package ai

type ServiceInterface interface {
	GenerateDialogueFromAI(request GenerateDialogueFromAIRequest) (*GenerateDialogueFromAIResponse, error)
}
