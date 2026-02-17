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

func (m *HskModel) GetWordsByHskSourceID(hskSourceID int) ([]Word, error) {
	var words []Word
	err := m.DB.NewSelect().
		Model(&words).
		Relation("WordTranslation").
		Relation("Example").
		Relation("Example.ExampleTranslation").
		Where("hsk_source_id = ?", hskSourceID).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (m *HskModel) GetWords(hskSourceID int, withPreviousLevel bool) ([]Word, error) {
	var words []Word
	query := m.DB.NewSelect().
		Model(&words).
		Join("JOIN hsk_sources ON word.hsk_source_id = hsk_sources.id")

	if withPreviousLevel {
		// Get words from level 1 to hskSourceID
		query = query.Where("hsk_sources.id <= ?", hskSourceID)
	} else {
		// Get words from specific level only
		query = query.Where("hsk_sources.id = ?", hskSourceID)
	}
	// Limit to 100 words
	query = query.Limit(100)
	query = query.OrderExpr("RANDOM()")

	err := query.Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return words, nil
}
