package services

import (
	"log"
	"oracao-bandas.com/src/database/resources/bands/entities"
	"oracao-bandas.com/src/database/resources/bands/repositories"
	"oracao-bandas.com/src/modules/bands/structs"
)

type SaveBandServiceInterface interface {
	SaveBand(band *structs.SaveBandDto) (error, entities.Band)
	LoadLastBands() []entities.Band
	SearchBands(page int, number int) []entities.Band
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

func (service *SaveBandService) LoadLastBands() []entities.Band {
	bands, err := service.repository.List(1, 3)

	if err != nil {
		log.Panic(err)
	}

	return bands
}

func (service *SaveBandService) SearchBands(page int, number int) []entities.Band {
	bands, err := service.repository.List(page, number)

	if err != nil {
		log.Panic(err)
	}

	return bands
}
