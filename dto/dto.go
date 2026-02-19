package dto

type GetWordsByHskSourceIDItem struct {
	ID         int                                  `json:"id"`
	Hanzi      string                               `json:"hanzi"`
	Pinyin     string                               `json:"pinyin"`
	English    string                               `json:"english_translation"`
	Indonesian string                               `json:"indonesian_translation"`
	Example    GetWordsByHskSourceIDResponseExample `json:"example"`
}

type GetWordsByHskSourceIDResponseExample struct {
	Hanzi      string `json:"hanzi"`
	Pinyin     string `json:"pinyin"`
	English    string `json:"english"`
	Indonesian string `json:"indonesian"`
}

type GetWordsByHskSourceIDResponse struct {
	List  []GetWordsByHskSourceIDItem `json:"list"`
	Total int                         `json:"total"`
}

type GenerateDialogueFromAIRequest struct {
	Complexity string `json:"complexity"`
}

type GenerateGradedTextRequest struct {
	Complexity int `json:"complexity"`
}
