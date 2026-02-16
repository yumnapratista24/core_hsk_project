package services

import (
	"core_hsk_project/dto"
	"core_hsk_project/model"
)

type Service struct {
	model model.ModelInterface
}

func NewService(model model.ModelInterface) ServiceInterface {
	return &Service{
		model: model,
	}
}

func (s *Service) GetWordsByHskSourceID(hskSourceID int) (dto.GetWordsByHskSourceIDResponse, error) {
	results, count, err := s.model.GetWordsByHskSourceID(hskSourceID)
	if err != nil {
		return dto.GetWordsByHskSourceIDResponse{}, err
	}

	response := buildGetWordsByHskSourceIDResponse(results, count)

	return response, nil
}

func buildGetWordsByHskSourceIDResponse(data []model.Word, count int) dto.GetWordsByHskSourceIDResponse {
	list := []dto.GetWordsByHskSourceIDItem{}
	for _, word := range data {
		list = append(list, dto.GetWordsByHskSourceIDItem{
			ID:         word.ID,
			Hanzi:      word.Hanzi,
			Pinyin:     word.Pinyin,
			English:    word.GetEnglish(),
			Indonesian: word.GetIndonesian(),
			Example: dto.GetWordsByHskSourceIDResponseExample{
				Hanzi:      word.Example.Hanzi,
				Pinyin:     word.Example.Pinyin,
				English:    word.Example.GetEnglish(),
				Indonesian: word.Example.GetIndonesian(),
			},
		})
	}
	return dto.GetWordsByHskSourceIDResponse{
		List:  list,
		Total: count,
	}
}
