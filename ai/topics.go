package ai

import (
	"math/rand"
	"time"
)

var (
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// HSK 1 Topics: The Basics (Level 1 Vocabulary)
// Focus: Survival, identification, and simple static descriptions.
var hsk1Topics = []string{
	"Self-Introduction",
	"Family Portrait",
	"Making Plans",
	"Telling Time",
	"At the Shop",
	"Transportation",
	"Location & Existence",
	"Weather (Basic)",
	"Likes & Dislikes",
	"Friendship",
}

// HSK 2 Topics: Daily Life (Level 2 Vocabulary)
// Focus: Routines, preferences, descriptions of action, and simple reasons.
var hsk2Topics = []string{
	"Ordering Food",
	"Talking About Hobbies",
	"Describing Appearance",
	"Weather Report",
	"Asking for Directions",
	"Being Sick",
	"Buying Clothes",
	"Work Routine",
	"Travel Experiences",
	"Comparisons",
}

// GetRandomTopic returns a random topic based on HSK level
func GetRandomTopic(hskLevel int) string {
	switch hskLevel {
	case 1:
		if len(hsk1Topics) > 0 {
			return hsk1Topics[rng.Intn(len(hsk1Topics))]
		}
	case 2:
		if len(hsk2Topics) > 0 {
			return hsk2Topics[rng.Intn(len(hsk2Topics))]
		}
	default:
		// For HSK 3+, mix both HSK 1 and 2 topics
		allTopics := append(hsk1Topics, hsk2Topics...)
		if len(allTopics) > 0 {
			return allTopics[rng.Intn(len(allTopics))]
		}
	}

	// Fallback topic if no topics available
	return "Daily Life: A simple story about everyday activities and experiences."
}

// GetAllTopics returns all available topics for reference
func GetAllTopics() map[int][]string {
	return map[int][]string{
		1: hsk1Topics,
		2: hsk2Topics,
	}
}
