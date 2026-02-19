package model

import (
	"time"

	"github.com/uptrace/bun"
)

type HskSource struct {
	bun.BaseModel `bun:"table:hsk_sources"`

	ID           int        `bun:"id,pk,autoincrement" json:"id"`
	Name         string     `bun:"name,notnull" json:"name"`
	Level        int        `bun:"level,notnull" json:"level"`
	YearReleased *int       `bun:"year_released" json:"year_released,omitempty"`
	HskVersion   *string    `bun:"hsk_version" json:"hsk_version,omitempty"`
	CreatedAt    time.Time  `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time  `bun:"updated_at,default:current_timestamp" json:"updated_at"`
	DeletedAt    *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at,omitempty"`
}

type Word struct {
	bun.BaseModel   `bun:"table:words"`
	WordTranslation []WordTranslation `bun:"rel:has-many,join:id=word_id"`
	Example         *Example          `bun:"rel:has-one,join:id=word_id"`
	HSKSource       *HskSource        `bun:"rel:belongs-to,join:hsk_source_id=id"`

	ID          int        `bun:"id,pk,autoincrement" json:"id"`
	Hanzi       string     `bun:"hanzi,notnull" json:"hanzi"`
	Pinyin      string     `bun:"pinyin,notnull" json:"pinyin"`
	HskSourceID *int       `bun:"hsk_source_id" json:"hsk_source_id,omitempty"`
	CreatedAt   time.Time  `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time  `bun:"updated_at,default:current_timestamp" json:"updated_at"`
	DeletedAt   *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at,omitempty"`
}

func (w *Word) GetEnglish() string {
	for _, translation := range w.WordTranslation {
		if translation.Language == "en" {
			return translation.Value
		}
	}
	return ""
}

func (w *Word) GetIndonesian() string {
	for _, translation := range w.WordTranslation {
		if translation.Language == "id" {
			return translation.Value
		}
	}
	return ""
}

type Example struct {
	bun.BaseModel      `bun:"table:examples"`
	ExampleTranslation []ExampleTranslation `bun:"rel:has-many,join:id=example_id"`

	ID        int        `bun:"id,pk,autoincrement" json:"id"`
	WordID    *int       `bun:"word_id" json:"word_id,omitempty"`
	Hanzi     string     `bun:"hanzi,notnull" json:"hanzi"`
	Pinyin    string     `bun:"pinyin,notnull" json:"pinyin"`
	CreatedAt time.Time  `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `bun:"updated_at,default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at,omitempty"`
}

func (e *Example) GetEnglish() string {
	for _, translation := range e.ExampleTranslation {
		if translation.Language == "en" {
			return translation.Value
		}
	}
	return ""
}

func (e *Example) GetIndonesian() string {
	for _, translation := range e.ExampleTranslation {
		if translation.Language == "id" {
			return translation.Value
		}
	}
	return ""
}

type WordTranslation struct {
	bun.BaseModel `bun:"table:word_translations"`

	ID        int        `bun:"id,pk,autoincrement" json:"id"`
	WordID    int        `bun:"word_id,notnull" json:"word_id"`
	Language  string     `bun:"language,notnull" json:"language"`
	Value     string     `bun:"value,notnull" json:"value"`
	CreatedAt time.Time  `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `bun:"updated_at,default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at,omitempty"`
}

type ExampleTranslation struct {
	bun.BaseModel `bun:"table:example_translations"`

	ID        int        `bun:"id,pk,autoincrement" json:"id"`
	ExampleID int        `bun:"example_id,notnull" json:"example_id"`
	Language  string     `bun:"language,notnull" json:"language"`
	Value     string     `bun:"value,notnull" json:"value"`
	CreatedAt time.Time  `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `bun:"updated_at,default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at,omitempty"`
}
