package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"oracao-bandas.com/src/database/entities"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (repository *AuthRepository) Register(name, email, password string) error {
	log.Printf("Registerig user for name: %s", name)

	var user entities.Auth
	user.Name = name
	user.Email = email
	user.Password = password

	response := repository.db.Create(&user)

	if response.Error != nil {
		log.Printf("Error trying to searching user. Details: %s", response.Error)
		err := errors.New("DB error query")
		return err
	}

	return nil
}

func (repository *AuthRepository) SearchByEmail(email string) (entities.Auth, error) {
	log.Printf("Searching user for email: %s", email)

	var user entities.Auth

	response := repository.db.First(&user, "email = ?", email)

	if response.Error != nil {
		log.Printf("Error trying to searching user. Details: %s", response.Error)
		err := errors.New("DB error query")
		return user, err
	}

	return user, nil
}
