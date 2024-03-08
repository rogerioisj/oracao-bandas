package memory

import (
	"log"
	"oracao-bandas.com/src/database/entities"
)

type DB struct {
	Bands []entities.Band
}

func Connector() *DB {
	log.Print("Success to initialize database connection. On memory database connected.")

	var db DB

	return &db
}
