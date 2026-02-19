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

func (m *HskModel) GetWords(hskSourceID int, withPreviousLevel bool) ([]Word, []Word, error) {
	var words []Word
	query := m.DB.NewSelect().
		Model(&words).
		Join("JOIN hsk_sources ON word.hsk_source_id = hsk_sources.id")
	query = query.Where("hsk_sources.id = ?", hskSourceID)
	query = query.Limit(100)
	query = query.OrderExpr("RANDOM()")
	err := query.Scan(context.Background())
	if err != nil {
		return nil, nil, err
	}

	var wordsPreviousLevel []Word
	if withPreviousLevel && hskSourceID > 1 {
		query := m.DB.NewSelect().
			Model(&wordsPreviousLevel).
			Relation("HSKSource")
		query = query.Where("hsk_source.id < ?", hskSourceID)
		query = query.Limit(100)
		query = query.OrderExpr("RANDOM()")
		err := query.Scan(context.Background())
		if err != nil {
			return nil, nil, err
		}
	}

	return words, wordsPreviousLevel, nil
}
