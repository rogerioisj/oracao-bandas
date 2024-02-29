package services

import (
	"oracao-bandas.com/src/api/modules/bands/structs"
	"oracao-bandas.com/src/database/resources/bands/entities"
	"oracao-bandas.com/src/database/resources/bands/repositories"
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

	//TODO: Inject connection

	/*response := postgres.Connect().Save(&savedBand)

	if response.Error != nil {
		return response.Error, entities.Band{}
	}*/

	err := service.repository.Save(&savedBand)

	if err != nil {
		return err, entities.Band{}
	}

	return nil, savedBand
}
