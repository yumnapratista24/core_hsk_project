package model

type ModelInterface interface {
	GetWordsByHskSourceID(hskSourceID int) ([]Word, error)
	GetWords(hskSourceID int, withPreviousLevel bool) ([]Word, []Word, error)
}
