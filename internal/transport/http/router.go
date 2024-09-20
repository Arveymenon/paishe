package http

import (
	"github.com/gin-gonic/gin"

	"github.com/Arveymenon/paishe/internal/config"
	"github.com/Arveymenon/paishe/internal/repository"
	service "github.com/Arveymenon/paishe/internal/services"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// var user_repository repository.UserRepository = repository.NewUserRepository(config.DB)
	user_repository := repository.NewUserRepository(config.DB)

	userService := service.NewUserService(user_repository)

	NewUserHandler(router, userService)

	return router
}
