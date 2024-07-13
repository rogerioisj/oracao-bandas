package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"oracao-bandas.com/src/configuration"
)

func Connect(config *configuration.Configuration) *gorm.DB {
	dsn := fmt.Sprintf("host='%s' user='%s' password='%s' dbname='%s' port='%v' sslmode=disable",
		config.Database.Host, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Fail to connect to database: %s", err)
		panic(err)
	}

	log.Printf("Using Postgres database at port %v", config.Database.Port)

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	return db
}
