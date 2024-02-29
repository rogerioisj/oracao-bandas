package main

import (
	"oracao-bandas.com/src/configuration"
	"oracao-bandas.com/src/database/orm"
	"oracao-bandas.com/src/database/postgres"
	"oracao-bandas.com/src/server"
)

func main() {
	config, _ := configuration.LoadEnvironmentVars()

	db := postgres.Connect(&config)

	orm.LoadEntities(db)

	server.StartServer(&config, db)
}
