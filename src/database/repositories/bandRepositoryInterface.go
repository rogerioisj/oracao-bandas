package repositories

import (
	"oracao-bandas.com/src/database/entities"
	"oracao-bandas.com/src/modules/bands/structs"
)

type BandRepositoryInterface interface {
	Save(dto *entities.Band) error
	Update(dto *structs.UpdateBandDto)
	Get(id string)
	List(page, number int, name string) ([]entities.Band, int, error)
}
