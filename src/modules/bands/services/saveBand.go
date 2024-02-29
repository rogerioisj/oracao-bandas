package services

import (
	"oracao-bandas.com/src/database/resources/bands/entities"
	"oracao-bandas.com/src/database/resources/bands/repositories"
	"oracao-bandas.com/src/modules/bands/structs"
)

type SaveBandServiceInterface interface {
	SaveBand(band *structs.SaveBandDto) (error, entities.Band)
}

type SaveBandService struct {
	repository repositories.BandRepositoryInterface
}

func NewSaveBandService(repository repositories.BandRepositoryInterface) *SaveBandService {
	return &SaveBandService{
		repository: repository,
	}
}

func (service *SaveBandService) SaveBand(band *structs.SaveBandDto) (error, entities.Band) {
	var savedBand entities.Band

	savedBand.Title = band.Title
	savedBand.Country = band.Country
	savedBand.Style = band.Style
	savedBand.Type = band.Type
	savedBand.UF = band.UF

	err := service.repository.Save(&savedBand)

	if err != nil {
		return err, entities.Band{}
	}

	return nil, savedBand
}
