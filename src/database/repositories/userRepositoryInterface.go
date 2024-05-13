package repositories

import (
	"oracao-bandas.com/src/database/entities"
)

type UserRepositoryInterface interface {
	SearchByLogin(login string) (entities.User, error)
	Register(name, login, password string) error
}
