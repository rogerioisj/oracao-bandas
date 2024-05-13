package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"oracao-bandas.com/src/database/entities"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (repository *SessionRepository) Save(login string) (entities.Session, error) {
	log.Printf("Registerig session for user: %s", login)

	var user entities.Session
	user.Login = login

	response := repository.db.Create(&user)

	if response.Error != nil {
		log.Printf("Error trying to save session. Details: %s", response.Error)
		err := errors.New("DB error query")
		return user, err
	}

	return user, nil
}

func (repository *SessionRepository) Delete(sessionId string) error {
	log.Printf("Searching session: %s", sessionId)

	_ = repository.db.Delete(&entities.Session{}, sessionId)

	return nil
}