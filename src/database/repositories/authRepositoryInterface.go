package repositories

import (
	"oracao-bandas.com/src/database/entities"
)

type AuthRepositoryInterface interface {
	SearchByLogin(email, password string) (entities.Auth, error)
	Register(name, email, password string) error
}
