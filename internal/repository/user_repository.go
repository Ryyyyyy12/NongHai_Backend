package repository

import (
	"backend/internal/domain/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user model.User) (newUser *model.User, err error)
	FindById(userId string) (user *model.User, err error)
	FindByIdWithPets(userId string) (user *model.User, err error)  // New method for preloading pets
	Update(user *model.User) error
}

type userRepository struct {
	DB *gorm.DB  // Changed to pointer to avoid potential issues with GORM
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		DB: db,
	}
}

// Create a new user in the database
func (r *userRepository) Create(user model.User) (newUser *model.User, err error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindById retrieves a user by their ID without preloading pets
func (r *userRepository) FindById(id string) (user *model.User, err error) {
	foundUser := new(model.User)
	if err := r.DB.First(&foundUser, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return foundUser, nil
}

// FindByIdWithPets retrieves a user by their ID and preloads pets from the pets table
func (r *userRepository) FindByIdWithPets(userId string) (user *model.User, err error) {
	foundUser := new(model.User)
	// Preload pets for the user
	if err := r.DB.Preload("Pets").First(&foundUser, "id = ?", userId).Error; err != nil {
		return nil, err
	}
	return foundUser, nil
}

// Update modifies the existing user in the database
func (r *userRepository) Update(user *model.User) error {
	return r.DB.Save(user).Error
}
