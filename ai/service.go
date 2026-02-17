package ai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
)

type Service struct {
	AIClient deepseek.Client
	Context  context.Context
}

func NewService(aiClient *deepseek.Client, ctx context.Context) ServiceInterface {
	return &Service{
		AIClient: *aiClient,
		Context:  ctx,
	}
}

func (s *Service) GenerateDialogueFromAI(req GenerateDialogueFromAIRequest) (*GenerateDialogueFromAIResponse, error) {
	respFormat := &request.ResponseFormat{
		Type: "json_object",
	}

	totalDialogue := GetDialogueFromComplexity(req.TextComplexity)
	complexity := getComplexityForLevel(req.TextComplexity)

	chatReq := &request.ChatCompletionsRequest{
		Model:          deepseek.DEEPSEEK_CHAT_MODEL,
		Stream:         false,
		ResponseFormat: respFormat,
		Messages: []*request.Message{
			{
				Role: "system",
				Content: fmt.Sprintf(`
					This is a system where we will generate some dialogue in Chinese, its pinyin, and its meaning in English.
					The goal of this system is to help the user to learn to read Chinese characters (Hanzi) by generating some dialogue about any topic.

					I will provide you some information regarding how many dialogue you need to make, and its complexity.
					There are 3 kinds of complexity, simple, medium, and complex. 
					For simple, maybe make it more straightforward similar to HSK books. Medium means a little bit more complex than simple, and complex means you need to make the dialogue more complex and similar to day-to-day Chinese people speaking.

					I will also provide you some words that you need to use in the dialogue. I will provide them below in a stringified format splitted with dash.

					These are the words that you might need to use based on each HSK levels : %s
					We are now using HSK level %d
				`, req.StringifiedWords, req.HSKLevel),
			},
			{
				Role: "system",
				Content: fmt.Sprintf(`
					[STRICT RULE] Total dialogue you need to generate is not more than %d.
					[STRICT RULE] The complexity of the dialogue is %s.
					
					You need to give the response in JSON format. There are 3 keys in the JSON format.
					{
						"dialogue": ["你好", "你好", "你叫什么名字？", "我叫张三", "你多大了？", "我二十五岁"],
						"pinyin": ["nǐ hǎo", "nǐ hǎo", "nǐ jiào shén me míng zì?", "wǒ jiào zhāng sān", "nǐ duō dà le?", "wǒ èr shí wǔ suì"],
						"english": ["hello", "hello", "what is your name?", "my name is Zhang San", "how old are you?", "I am 25 years old"]
						"error": null
					}

					The "dialogue" is an array of string, where each string is a dialogue in Chinese. The dialogue will only be between 2 people, after each dialogue the other person will answer in the next dialogue.
					The "pinyin" is an array of string, where each string is a pinyin of the dialogue in Chinese. The pinyin will be in the same order as the dialogue.
					The "english" is an array of string, where each string is a translation of the dialogue in English. The translation will be in the same order as the dialogue.
					
					EXAMPLE JSON OUTPUT [SUCCESS CASE]:
					{
						"dialogue": ["你好", "你好", "你叫什么名字？", "我叫张三", "你多大了？", "我二十五岁"],
						"pinyin": ["nǐ hǎo", "nǐ hǎo", "nǐ jiào shén me míng zì?", "wǒ jiào zhāng sān", "nǐ duō dà le?", "wǒ èr shí wǔ suì"],
						"english": ["hello", "hello", "what is your name?", "my name is Zhang San", "how old are you?", "I am 25 years old"]
						"error": null
					}

					EXAMPLE JSON OUTPUT [ERROR CASE]:
					{
						"dialogue": null,
						"pinyin": null,
						"english": null,
						"error": "The topic is not valid, please choose another topic"
					}
				`, totalDialogue, complexity),
			},
		},
	}

	response, err := s.AIClient.CallChatCompletionsChat(s.Context, chatReq)
	if err != nil {
		return nil, err
	}

	jsonString := response.Choices[0].Message.Content

	var result GenerateDialogueFromAIResponse
	err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %v", err)
	}

	// Check for AI error
	if result.Error != nil && *result.Error != "" {
		return nil, fmt.Errorf("AI error: %s", *result.Error)
	}

	// Validate response
	if len(result.Dialogue) == 0 {
		return nil, fmt.Errorf("no dialogue generated")
	}

	// Check array lengths match
	if len(result.Dialogue) != len(result.Pinyin) || len(result.Dialogue) != len(result.English) {
		return nil, fmt.Errorf("array length mismatch in AI response")
	}

	return &result, nil
}
