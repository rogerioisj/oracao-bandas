package repositories

import (
	"oracao-bandas.com/src/database/entities"
)

type SessionRepositoryInterface interface {
	Save(login string) (entities.Session, error)
	Delete(sessionId string) error
}
