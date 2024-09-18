package service

import "errors"

type User struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}

type UserService interface {
    GetUser(id uint) (*User, error)
}

type userService struct {
    users map[uint]*User
}

func NewUserService() UserService {
    return &userService{
        users: map[uint]*User{
            1: {ID: 1, Name: "John Doe"},
        },
    }
}

func (s *userService) GetUser(id uint) (*User, error) {
    user, exists := s.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}
