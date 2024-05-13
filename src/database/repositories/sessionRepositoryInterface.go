package repositories

import (
	"oracao-bandas.com/src/database/entities"
)

type SessionRepositoryInterface interface {
	Save(login string) (entities.User, error)
	Delete(login string) error
}
