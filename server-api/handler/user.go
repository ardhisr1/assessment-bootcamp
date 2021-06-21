package handler

import "github.com/gin-gonic/gin"

type UserHandler struct {
	service user.Service
}

func NewHandler(service user.Service) *UserHandler {
	return &UserHandler{service}


