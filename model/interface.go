package model

type ModelInterface interface {
	GetWordsByHskSourceID(hskSourceID int) ([]Word, int, error)
}
