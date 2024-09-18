package http

import (
    "github.com/gin-gonic/gin"
    "github.com/Arveymenon/paishe/internal/services"
    "net/http"
    "strconv"
)

type UserHandler struct {
    service service.UserService
}

func NewUserHandler(router *gin.Engine, service service.UserService) {
    handler := &UserHandler{service: service}
    router.GET("/users/:id", handler.GetUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
    id := c.Param("id")
    idUint, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    user, err := h.service.GetUser(uint(idUint))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}