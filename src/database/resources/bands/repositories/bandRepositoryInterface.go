package repositories

import (
	"oracao-bandas.com/src/database/resources/bands/entities"
	"oracao-bandas.com/src/modules/bands/structs"
)

type BandRepositoryInterface interface {
	Save(dto *structs.SaveBandDto) entities.Band
	Update(dto *structs.UpdateBandDto) entities.Band
	Get(id string) entities.Band
	List(page int, number int) []entities.Band
}
