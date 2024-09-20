package http

import (
	service "github.com/Arveymenon/paishe/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	userService := service.NewUserService(db)

	NewUserHandler(router, userService)

	return router
}
