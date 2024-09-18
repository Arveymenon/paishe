package http

import (
    "github.com/gin-gonic/gin"
    "github.com/Arveymenon/paishe/internal/services"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    userService := service.NewUserService()
    
    NewUserHandler(router, userService)
    
    return router
}