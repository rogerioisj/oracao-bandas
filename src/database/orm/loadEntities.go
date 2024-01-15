package orm

import (
	"gorm.io/gorm"
	"log"
	"oracao-bandas.com/src/database/resources/bands/entities"
)

func LoadEntities(database *gorm.DB) {
	err := database.AutoMigrate(&entities.Band{})
	if err != nil {
		log.Fatalf("Error to load entities %s", err)
	}
}
