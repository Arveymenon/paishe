package service

import (
	"errors"
	"log"

	"github.com/Arveymenon/paishe/internal/domain"
	"github.com/Arveymenon/paishe/internal/repository"
	"gorm.io/gorm"
)

var userRepository repository.UserRepository

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserService interface {
	GetUser(id uint) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
}

type userService struct {
	users map[uint]*User
}

func NewUserService(db *gorm.DB) UserService {
	userRepository = repository.NewUserRepository(db)

	return &userService{
		users: map[uint]*User{
			1: {ID: 1, Name: "John Doe"},
		},
	}
}

func (s *userService) GetUser(id uint) (*domain.User, error) {
	user, err := userRepository.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) CreateUser(user *domain.User) (*domain.User, error) {
	log.Print("Create new user log")
	response, err := userRepository.CreateUser(user)
	log.Print("Create new user succeded - %s", err)
	if err != nil {
		return nil, err
	}
	log.Print(response)
	return response, nil
}
