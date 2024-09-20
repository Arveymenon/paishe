package service

import (
	"log"

	"github.com/Arveymenon/paishe/internal/domain"
	"github.com/Arveymenon/paishe/internal/repository"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserService interface {
	GetUser(id uint) (*domain.User, error)
	CreateNewUser(user domain.User) (*domain.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) userService {
	s := userService{}
	s.userRepository = userRepository
	return s
}

func (s *userService) GetUser(id uint) (*domain.User, error) {
	user, err := s.userRepository.GetUserByID(id) // Call the actual repository method
	if err != nil {
		return nil, err // Return the error from the repository
	}
	return user, nil
}

func (s userService) CreateNewUser(reqUser domain.User) (*domain.User, error) {
	log.Print("Create new user log")

	response, err := s.userRepository.CreateUser(&reqUser)
	log.Print("Create new user succeded - %s", err)
	if err != nil {
		return nil, err // Return the error from the repository
	}
	log.Print(response)
	return response, nil
}
