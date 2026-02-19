package services

import (
	"core_hsk_project/dto"
	"core_hsk_project/model"
)

type ServiceInterface interface {
	GetWordsByHskSourceID(hskSourceID int) (dto.GetWordsByHskSourceIDResponse, error)
	GetWordsWithPreviousLevel(hskSourceID int) ([]model.Word, []model.Word, error)
}
