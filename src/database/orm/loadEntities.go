package orm

import (
	"gorm.io/gorm"
	"log"
	"oracao-bandas.com/src/database/entities"
)

func LoadEntities(database *gorm.DB) {
	err := database.AutoMigrate(&entities.Band{}, &entities.Auth{})
	if err != nil {
		log.Fatalf("Error to load entities %s", err)
	}
}
