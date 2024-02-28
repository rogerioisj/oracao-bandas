package repositories

import (
	"oracao-bandas.com/src/api/modules/bands/structs"
	"oracao-bandas.com/src/database/resources/bands/entities"
)

type BandRepositoryInterface interface {
	Save(dto *structs.SaveBandDto) entities.Band
	Update(dto *structs.UpdateBandDto) entities.Band
	Get(id string) entities.Band
	List(page int, number int) []entities.Band
}
