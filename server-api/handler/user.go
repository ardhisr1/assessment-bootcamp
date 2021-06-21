package handler

import (
	"server-api/entity"
	"server-api/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewHandler(service user.Service) *userHandler {
	return &userHandler{service}

func (h *userHandler) RegisterUser(c *gin.Context) {
	 var input entity.UserInput

	 newUser, err := h.service.SaveNewUser(input)
	 
	 response := {
		 201,
		 "success",
	 }
	 c.JSON(201, response)
}
