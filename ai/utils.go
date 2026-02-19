package ai

import "strings"

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

// ConjunctionsWhitelist contains essential glue words to allow compound sentences.
var conjunctionsWhitelist = []string{
	"和",  // and
	"因为", // because
	"所以", // therefore
	"但是", // but
	"可是", // but/however
	"然后", // then
	"如果", // if
	"或者", // or (statements)
	"还是", // or (questions)
	"不但", // not only
	"而且", // but also
}

// NamesWhitelist contains standard names for story characters.
var namesWhitelist = []string{
	"李华", // Li Hua
	"王明", // Wang Ming
	"张雪", // Zhang Xue
	"大卫", // David
	"玛丽", // Mary
}

// LocationsWhitelist gives your characters places to go.
var locationsWhitelist = []string{
	"北京", // Beijing
	"上海", // Shanghai
	"沈阳", // Shenyang
}

func getWhitelistedWords() []string {
	return append(conjunctionsWhitelist, append(namesWhitelist, locationsWhitelist...)...)
}

func GetStringifiedWhitelistedWords() string {
	var stringifiedWords strings.Builder
	for _, word := range getWhitelistedWords() {
		stringifiedWords.WriteString(word + "-")
	}

	return stringifiedWords.String()
}
