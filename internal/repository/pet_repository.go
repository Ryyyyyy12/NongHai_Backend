package repository

import (
	"backend/internal/domain/model"
	"gorm.io/gorm"
)

type IPetRepository interface {
	FindById(petId string) (user *model.Pet, err error)
}

type petRepository struct {
	DB gorm.DB
}

func NewPetRepository(db gorm.DB) IPetRepository {
	return &petRepository{
		DB: db,
	}
}

func (r *petRepository) FindById(id string) (pet *model.Pet, err error) {
	foundPet := new(model.Pet)
	if err := r.DB.First(&foundPet, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return foundPet, nil
}
