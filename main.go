package main

import (
	"oracao-bandas.com/src/database/orm"
	"oracao-bandas.com/src/database/postgres"
	"oracao-bandas.com/src/server"
)

func main() {
	db := postgres.Connect()

	orm.LoadEntities(db)

	server.StartServer()
}
