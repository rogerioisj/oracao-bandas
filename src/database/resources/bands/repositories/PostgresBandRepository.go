package repositories

import (
	"gorm.io/gorm"
	"log"
	"oracao-bandas.com/src/api/modules/bands/structs"
	"oracao-bandas.com/src/database/resources/bands/entities"
)

type PostgresBandRepository struct {
	db *gorm.DB
}

func NewPostgresBandRepository(db *gorm.DB) *PostgresBandRepository {
	return &PostgresBandRepository{
		db: db,
	}
}

func (repository *PostgresBandRepository) Save(band *entities.Band) error {
	log.Printf("Saving band: %s", band)
	repository.db.Create(band)

	return nil
}

func (repository *PostgresBandRepository) Update(dto *structs.UpdateBandDto) {

}

func (repository *PostgresBandRepository) Get(id string) {

}

func (repository *PostgresBandRepository) List(page int, number int) {

}
