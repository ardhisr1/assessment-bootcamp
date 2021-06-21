package user

import (
	"server-api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (entity.User, error)
	FindByID(userID string) (entity.User, error)
	Create() (entity.User, error)
	UpdateByID(userID string) (entity.User, error)
	DeleteByID(userID string) (interface{}, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
