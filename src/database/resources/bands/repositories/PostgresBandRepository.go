package repositories

import (
	"gorm.io/gorm"
	"log"
	"oracao-bandas.com/src/database/resources/bands/entities"
	"oracao-bandas.com/src/modules/bands/structs"
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

func (repository *PostgresBandRepository) List(page int, number int) ([]entities.Band, error) {
	log.Printf("Listing bands for page: %v and number: %v", page, number)

	var bands []entities.Band

	var offset int

	if page < 1 {
		offset = 0
	} else {
		offset = (page - 1) * number
	}

	response := repository.db.Order("created_at desc").Order("title").Offset(offset).Limit(number).Find(&bands)

	if response.Error != nil {
		log.Printf("Error to load Bands. Details: %s", response.Error)
		return nil, response.Error
	}

	return bands, nil
}
