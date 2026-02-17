package ai

type GenerateDialogueFromAIRequest struct {
	StringifiedWords string
	TextComplexity   int
	HSKLevel         int
}

type GenerateDialogueFromAIResponse struct {
	Dialogue []string `json:"dialogue"`
	Pinyin   []string `json:"pinyin"`
	English  []string `json:"english"`
	Error    *string  `json:"error"`
}
