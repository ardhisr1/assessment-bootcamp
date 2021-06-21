package handler

import (
	"server-api/auth"
	"server-api/entity"
	"server-api/helper"
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
		responseError := helper.FailResponse(500, "internal server error", err.Error())
		c.JSON(500, responseError)
	}
	response := helper.APIResponse(201, "create user succes", newUser)
	c.JSON(201, response)
}

func (h *userHandler) GetUserDetailHandler(c *gin.Context) {
	var currentUser = string(int(c.MustGet("currentUser").(int)))

	user, err := h.service.GetUserByID(currentUser)

	if err != nil {
		responseError := helper.FailResponse(500, "internal server error", err.Error())
		c.JSON(500, responseError)
	}
	response := helper.APIResponse(201, "get user succes", user)
	c.JSON(201, response)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var input entity.UserLogin

	user, err := h.service.Login(input)

	if err != nil {
		responseError := helper.FailResponse(500, "internal server error", err.Error())
		c.JSON(500, responseError)
	}

	token, err := h.authService.GenerateToken(user.ID)

	if err != nil {
		responseError := helper.FailResponse(500, "internal server error", err.Error())
		c.JSON(500, responseError)
	}

	response := helper.APIResponse(201, "create user succes", gin.H{"Authorization": token})
	c.JSON(201, response)
}
