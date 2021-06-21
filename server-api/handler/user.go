package handler

import (
	"server-api/auth"
	"server-api/entity"
	"server-api/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     user.Service
	authService auth.Service
}

func NewUserHandler(service user.Service, authService auth.Service) *userHandler {
	return &userHandler{service, authService}
}

func (h *userHandler) RegisterUserHandler(c *gin.Context) {
	var input entity.UserInput

	newUser, err := h.service.SaveNewUser(input)

	if err != nil {
	}
	c.JSON(201, newUser)
}
