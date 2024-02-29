package repositories

import (
	"oracao-bandas.com/src/api/modules/bands/structs"
	"oracao-bandas.com/src/database/resources/bands/entities"
)

type BandRepositoryInterface interface {
	Save(dto *entities.Band) error
	Update(dto *structs.UpdateBandDto)
	Get(id string)
	List(page int, number int)
}
