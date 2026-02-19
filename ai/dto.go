package ai

type WordItem struct {
	Hanzi    string
	Pinyin   string
	English  string
	HSKLevel int
}

type GenerateDialogueFromAIRequest struct {
	StringifiedWords   string
	PreviousLevelWords string
	TextComplexity     int
	HSKLevel           int
}
type GenerateGradedTextFromAIRequest struct {
	Words          []WordItem
	PrevLevelWords []WordItem
	TextComplexity int
	HSKLevel       int
	Topic          string
}

type GenerateDialogueFromAIResponse struct {
	Dialogue []string `json:"dialogue"`
	Pinyin   []string `json:"pinyin"`
	English  []string `json:"english"`
	Error    *string  `json:"error"`
}

type LineDetailsItem struct {
	Word    string `json:"word"`
	Pinyin  string `json:"pinyin"`
	English string `json:"english"`
}

type GenerateGradedTextFromAIResponse struct {
	Title       string            `json:"title"`
	LineDetails []LineDetailsItem `json:"line_details"`
	English     []string          `json:"english"`
	Error       *string           `json:"error"`
}
