package repository

import (
	"backend/internal/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user model.User) (newUser *model.User, err error)
	FindById(userId string) (user *model.User, err error)
}


type userRepository struct {
	DB gorm.DB
}

func NewUserRepository(db gorm.DB) IUserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Create(user model.User) (newUser *model.User, err error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindById(id string) (user *model.User, err error) {
	foundUser := new(model.User)
	if err := r.DB.First(&foundUser, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return foundUser, nil
}
