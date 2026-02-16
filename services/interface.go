package services

import "core_hsk_project/dto"

type ServiceInterface interface {
	GetWordsByHskSourceID(hskSourceID int) (dto.GetWordsByHskSourceIDResponse, error)
}
