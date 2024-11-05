package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"oracao-bandas.com/src/database/entities"
	"oracao-bandas.com/src/database/repositories"
	auth "oracao-bandas.com/src/modules/auth/structs"
)

type AuthServiceInterface interface {
	CreateUser(dto *auth.CreateUserDto) error
	UpdateUser(dto *auth.CreateUserDto, login string) error
	Login(dto *auth.LoginDto) (string, error)
	Logout(sessionId string)
	ListUsers(page, number int, name string) ([]entities.User, int)
	GetUserByLogin(login string) (entities.User, error)
	GetUserBySession(sessionId string) string
}

type AuthService struct {
	userRepository    repositories.UserRepositoryInterface
	sessionRepository repositories.SessionRepositoryInterface
}

func NewAuthService(userRepository repositories.UserRepositoryInterface, sessionRepository repositories.SessionRepositoryInterface) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (service *AuthService) GetUserByLogin(login string) (entities.User, error) {
	user, errSearch := service.userRepository.SearchByLogin(login)

	if errSearch != nil {
		log.Printf("Error to search user %s", errSearch)
		return user, errors.New("error to create user")
	}

	return user, nil
}

func (service *AuthService) UpdateUser(dto *auth.CreateUserDto, login string) error {
	user, errSearch := service.userRepository.SearchByLogin(login)

	if errSearch != nil {
		log.Printf("Error to search user %s", errSearch)
		return errors.New("an error occurs trying to update user")
	}

	if user.Login == "" {
		log.Printf("User not found for %s", dto.Login)
		return errors.New("user not found")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error to hash password %s", err)
		return errors.New("an error occurs trying to update user")
	}

	err = service.userRepository.Update(login, dto.Name, dto.Login, string(hashedPassword))

	if err != nil {
		log.Printf("Error to update user %s", err)

		return errors.New("an error occurs trying to update user")
	}

	return nil
}

func (service *AuthService) CreateUser(dto *auth.CreateUserDto) error {
	user, errSearch := service.userRepository.SearchByLogin(dto.Login)

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

	err = service.userRepository.Register(dto.Name, dto.Login, string(hashedPassword))

	if err != nil {
		log.Printf("Error to create user %s", err)

		return errors.New("error to create user")
	}

	return nil
}

func (service *AuthService) Login(dto *auth.LoginDto) (string, error) {
	user, err := service.userRepository.SearchByLogin(dto.Login)
	if err != nil {
		log.Print("User not found")
		return "", err
	}

	checkErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

	if checkErr != nil {
		log.Printf("Wrong password for user %s", user.Login)
		return "", errors.New("Wrong password")
	}

	session, err := service.sessionRepository.Save(user.Login)

	if err != nil {
		log.Printf("Error to save session for %s", user.Login)
		return "", errors.New("Error to save session")
	}

	parsedId := session.ID.String()

	return parsedId, nil
}

func (service *AuthService) Logout(sessionId string) {
	_ = service.sessionRepository.Delete(sessionId)

	return
}

func (service *AuthService) ListUsers(page int, number int, name string) ([]entities.User, int) {
	users, total, err := service.userRepository.List(page, number, name)
	if err != nil {
		log.Print(err)
	}

	return users, total
}

func (service *AuthService) GetUserBySession(sessionId string) string {
	session, _ := service.sessionRepository.Search(sessionId)

	return session.Login
}
