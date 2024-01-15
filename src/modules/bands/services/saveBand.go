package services

import (
	"oracao-bandas.com/src/database/postgres"
	"oracao-bandas.com/src/database/resources/bands/entities"
	"oracao-bandas.com/src/modules/bands/structs"
)

func SaveBand(band *structs.SaveBandDto) (error, entities.Band) {
	var savedBand entities.Band

	savedBand.Title = band.Title
	savedBand.Country = band.Country
	savedBand.Style = band.Style
	savedBand.Type = band.Type
	savedBand.UF = band.UF

	response := postgres.Connect().Save(&savedBand)

	if response.Error != nil {
		return response.Error, entities.Band{}
	}

	return nil, savedBand
}
