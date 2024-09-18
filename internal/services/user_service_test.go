package service

import (
    "testing"
    "github.com/Arveymenon/paisheinternal/domain"
    "github.com/Arveymenon/paisheinternal/repository"
    "github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
    repo := &repository.MockUserRepository{}
    service := NewUserService(repo)

    user := &domain.User{ID: 1, Name: "John Doe"}
    repo.On("GetUserByID", uint(1)).Return(user, nil)

    result, err := service.GetUser(1)
    assert.NoError(t, err)
    assert.Equal(t, "John Doe", result.Name)
}
