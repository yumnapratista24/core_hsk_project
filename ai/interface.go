package ai

type ServiceInterface interface {
	GenerateDialogueFromAI(request GenerateDialogueFromAIRequest) (*GenerateDialogueFromAIResponse, error)
	GenerateGradedTextFromAI(req GenerateGradedTextFromAIRequest) (*GenerateGradedTextFromAIResponse, error)
}
