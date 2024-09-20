package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Arveymenon/paishe/internal/domain"
	service "github.com/Arveymenon/paishe/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(router *gin.Engine, service service.UserService) {
	handler := &UserHandler{service: service}
	router.GET("/users/:id", handler.GetUser)
	router.POST("/users/create", handler.CreateUser)
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

func (h *UserHandler) CreateUser(c *gin.Context) {
	var payload struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	log.Printf("Payload username - %s", payload.Name)

	user := domain.User{
		Name:  payload.Name,
		Email: payload.Email,
		Phone: payload.Phone,
	}
	h.service.CreateNewUser(user)

	c.JSON(http.StatusOK, gin.H{"success": user})
}
