package ai

func getComplexityForLevel(level int) string {
	switch level {
	case 1:
		return SIMPLE_COMPLEXITY
	case 2:
		return MEDIUM_COMPLEXITY
	case 3:
		return COMPLEX_COMPLEXITY
	default:
		return SIMPLE_COMPLEXITY
	}
}

const (
	SIMPLE_COMPLEXITY  = "simple"
	MEDIUM_COMPLEXITY  = "medium"
	COMPLEX_COMPLEXITY = "complex"

	SIMPLE_TOTAL_DIALOGUE  = 12
	MEDIUM_TOTAL_DIALOGUE  = 16
	COMPLEX_TOTAL_DIALOGUE = 20
)

func GetDialogueFromComplexity(complexity int) int {
	switch complexity {
	case 1:
		return SIMPLE_TOTAL_DIALOGUE
	case 2:
		return MEDIUM_TOTAL_DIALOGUE
	case 3:
		return COMPLEX_TOTAL_DIALOGUE
	default:
		return SIMPLE_TOTAL_DIALOGUE
	}
}
