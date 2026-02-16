package model

import (
	"context"

	"github.com/uptrace/bun"
)

type HskModel struct {
	DB *bun.DB
}

func NewHskModel(db *bun.DB) ModelInterface {
	return &HskModel{DB: db}
}

func (m *HskModel) GetWordsByHskSourceID(hskSourceID int) ([]Word, int, error) {
	var words []Word
	count, err := m.DB.NewSelect().
		Model(&words).
		Relation("WordTranslation").
		Relation("Example").
		Relation("Example.ExampleTranslation").
		Where("hsk_source_id = ?", hskSourceID).
		ScanAndCount(context.Background())
	if err != nil {
		return nil, 0, err
	}
	return words, count, nil
}
