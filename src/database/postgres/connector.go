package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=admin password=admin dbname=bands port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Fail to connect to database: %s", err)
	}

	log.Printf("Using Postgres database at port 5432 and Bands database")

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	return db
}
