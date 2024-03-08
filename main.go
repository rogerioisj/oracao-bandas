package main

import (
	"oracao-bandas.com/src/configuration"
	"oracao-bandas.com/src/database/orm"
	"oracao-bandas.com/src/database/postgres"
	"oracao-bandas.com/src/server"
)

func main() {
	config, err := configuration.LoadEnvironmentVars()

	if err != nil {
		panic(err)
	}

	db := postgres.Connect(&config)

	orm.LoadEntities(db)

	server.StartServer(&config, db)
}
