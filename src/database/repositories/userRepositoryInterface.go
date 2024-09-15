package repositories

import (
	"oracao-bandas.com/src/database/entities"
)

type UserRepositoryInterface interface {
	SearchByLogin(login string) (entities.User, error)
	Register(name, login, password string) error
	Update(originLogin, name, login, password string) error
	List(page, number int, name string) ([]entities.User, int, error)
}
