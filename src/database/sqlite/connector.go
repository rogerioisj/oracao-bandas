package sqlite

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error to initialize database %s", err)
		log.Fatal(err)
	}

	log.Print("Success to initialize database connection. SQLite database connected")

	return db
}
