package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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

func (s *Service) GenerateGradedTextFromAI(req GenerateGradedTextFromAIRequest) (*GenerateGradedTextFromAIResponse, error) {
	respFormat := &request.ResponseFormat{
		Type: "json_object",
	}

	var stringifiedWords strings.Builder
	for _, word := range req.Words {
		stringifiedWords.WriteString(word.Hanzi + "-")
	}

	var stringifiedPrevWords strings.Builder
	for _, word := range req.PrevLevelWords {
		stringifiedPrevWords.WriteString(word.Hanzi + "-")
	}

	complexity := getComplexityForLevel(req.TextComplexity)
	topic := req.Topic

	chatReq := &request.ChatCompletionsRequest{
		Model:          deepseek.DEEPSEEK_CHAT_MODEL,
		Stream:         false,
		ResponseFormat: respFormat,
		Messages: []*request.Message{
			{
				Role: "system",
				Content: fmt.Sprintf(`
					This is a system where we will generate some graded text in Chinese, its pinyin, and its meaning in English. Graded text here means like a short-story and we avoid to use dialogue format here.
					The goal of this system is to help the user to learn to read Chinese characters (Hanzi) by generating some graded text about any topic.

					There are 3 kinds of complexity, simple, medium, and complex. 
					For simple, maybe make it more straightforward similar to HSK books. Medium means a little bit more complex than simple, and complex means you need to make the dialogue more complex.

					I will also provide you some words that you need to use in the graded text. I will provide them below in a stringified format splitted with dash.

					[STRICT RULE] These are the words that you CAN use from selected HSK level: %s
					[STRICT RULE] These are the words that you CAN use from previous HSK level: %s
					[STRICT RULE] You are allowed to use these whitelisted words as well to help create better sentences: %s
					[STRICT RULE] DON'T USE OTHER WORDS THAT NOT IN THE PROVIDED LIST
					We are now using HSK level %d
					We are now using topic %s
				`, stringifiedWords.String(), GetStringifiedWhitelistedWords(), stringifiedPrevWords.String(), req.HSKLevel, topic),
			},
			{
				Role: "system",
				Content: fmt.Sprintf(`
					[STRICT RULE] The complexity of the dialogue is %s.
					[STRICT RULE] While we are not expecting you to return dialogue format, please split the text to multiple lines per SENTENCES.
					[STRICT RULE] Don't use any dialogue format in the graded text.
					[STRICT RULE] CHARACTER LIMITS BY COMPLEXITY:
					- Simple: 50-120 characters.
					- Medium: 200-350 characters.
					- Complex: 500-700 characters.
					
					You need to give the response in JSON format. There are 4 keys in the JSON format.
					{
						"error": null,
						"title": "title",
						"english": ["hello"],
						"line_details": [
							{
								word: "你好",
								pinyin: "nǐhǎo",
								english: "hello",
							},
							...
						]
					}

					The "line_details" is an array of object, where each object is a word. You need to prepare this too.
					and error.
					The "english" is an array of single or multi line text in english. This field represents the meaning based on context of the sentences/text.
					The "title" is a short title for the graded text with format 'Hanzi (english translation)'.

					Since we are using topic, if the topic seems not valid, please return error message. The definition of not valid is not appropriate in public in terms of ethical.
					
					EXAMPLE JSON OUTPUT [SUCCESS CASE]:
					{
						"error": null,
						"title": "title",
						"line_details": [
							{
								word: "你好",
								pinyin: "nǐhǎo",
								english: "hello",
							},
							...
						]
					}

					EXAMPLE JSON OUTPUT [ERROR CASE]:
					{	
						"title": "title",
						"error": "The topic is not valid, please choose another topic",
						"english": null,
						"line_details": null
					}
				`, complexity),
			},
		},
	}

	response, err := s.AIClient.CallChatCompletionsChat(s.Context, chatReq)
	if err != nil {
		return nil, err
	}

	jsonString := response.Choices[0].Message.Content

	var result GenerateGradedTextFromAIResponse
	err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %v", err)
	}

	// Check for AI error
	if result.Error != nil && *result.Error != "" {
		return nil, fmt.Errorf("AI error: %s", *result.Error)
	}

	return &result, nil
}
