package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"oracao-bandas.com/src/database/repositories"
	auth "oracao-bandas.com/src/modules/auth/structs"
)

type AuthServiceInterface interface {
	CreateUser(dto *auth.CreateUserDto) error
	Login(dto *auth.LoginDto) error
	Logout()
}

type AuthService struct {
	repository repositories.UserRepositoryInterface
}

func NewAuthService(repository repositories.UserRepositoryInterface) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (service *AuthService) CreateUser(dto *auth.CreateUserDto) error {
	user, errSearch := service.repository.SearchByLogin(dto.Login)

	if errSearch != nil {
		log.Printf("Error to search user %s", errSearch)
		return errors.New("error to create user")
	}

	if user.Login != "" {
		log.Printf("User already registered for %s", dto.Login)
		return errors.New("user already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error to hash password %s", err)
		return errors.New("error to create user")
	}

	err = service.repository.Register(dto.Name, dto.Login, string(hashedPassword))

	if err != nil {
		log.Printf("Error to create user %s", err)

		return errors.New("error to create user")
	}

	return nil
}

func (service *AuthService) Login(dto *auth.LoginDto) error {
	user, err := service.repository.SearchByLogin(dto.Login)
	if err != nil {
		log.Print("User not found")
		return err
	}

	checkErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

	if checkErr != nil {
		log.Printf("Wrong password for user %s", user.Login)
		return errors.New("Wrong password")
	}

	return nil
}

func (service *AuthService) Logout() {
	//TODO implement me
	panic("implement me")
}
