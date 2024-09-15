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

func (repository *UserRepository) List(page, number int, name string) ([]entities.User, int, error) {
	log.Printf("Listing users for page: %v and number: %v", page, number)

	var users []entities.User
	var total int64

	var offset int

	if page < 1 {
		offset = 0
	} else {
		offset = (page - 1) * number
	}

	response := repository.db.Order("created_at desc").Offset(offset).Limit(number).Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").Find(&users).Count(&total)
	_ = repository.db.Model(entities.User{}).Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").Count(&total)

	if response.Error != nil {
		log.Printf("Error to load Users. Details: %s", response.Error)
		return nil, 0, response.Error
	}

	newTotal := int(total)

	return users, newTotal, nil
}
