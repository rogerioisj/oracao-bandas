package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"oracao-bandas.com/src/database/entities"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repository *UserRepository) Register(name, login, password string) error {
	log.Printf("Registerig user for name: %s", name)

	var user entities.User
	user.Name = name
	user.Login = login
	user.Password = password

	response := repository.db.Create(&user)

	if response.Error != nil {
		log.Printf("Error trying to searching user. Details: %s", response.Error)
		err := errors.New("DB error query")
		return err
	}

	return nil
}

func (repository *UserRepository) SearchByLogin(email string) (entities.User, error) {
	log.Printf("Searching user for email: %s", email)

	var user entities.User

	response := repository.db.Limit(1).First(&user, "login = ?", email)

	if response.Error != nil {

		if response.Error.Error() == "record not found" {
			return user, nil
		}

		log.Printf("Error trying to searching user. Details: %s", response.Error)
		err := errors.New("DB error query")
		return user, err
	}

	return user, nil
}
