package user

import (
	"server-api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.User, error)
	FindByID(userID string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
	UpdateByID(userID string, data map[string]interface{}) (entity.User, error)
	DeleteByID(userID string) (interface{}, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindByID(UserID string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", UserID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Create(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateByID(UserID string, data map[string]interface{}) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", UserID).Find(&user).Error; err != nil {
		return user, err
	}

	if err := r.db.Model(&user).Where("id = ?", UserID).Updates(data).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteByID(UserID string) (interface{}, error) {
	if err := r.db.Where("id = ?", UserID).Delete(&entity.User{}).Error; err != nil {
		return "error", err
	}

	status := "user delete success"
	return status, nil
}
