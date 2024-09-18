package repository

import (
    "github.com/Arveymenon/paishe/internal/domain"
    "gorm.io/gorm"
)

type UserRepository interface {
    GetUserByID(id uint) (*domain.User, error)
    CreateUser(user *domain.User) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(id uint) (*domain.User, error) {
    var user domain.User
    if err := r.db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) CreateUser(user *domain.User) error {
    return r.db.Create(user).Error
}
