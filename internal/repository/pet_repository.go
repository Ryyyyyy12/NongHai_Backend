package repository

import (
	"backend/internal/domain/model"

	"gorm.io/gorm"
)

type IPetRepository interface {
	FindById(petId string) (*model.Pet, error)
	Create(pet *model.Pet) (*model.Pet, error)
	Update(petId string, updateData map[string]interface{}) (*model.Pet, error)
}

type petRepository struct {
	DB *gorm.DB
}

func NewPetRepository(db *gorm.DB) IPetRepository {
	return &petRepository{
		DB: db,
	}
}

// FindById retrieves a pet by ID
func (r *petRepository) FindById(id string) (*model.Pet, error) {
	foundPet := new(model.Pet)
	if err := r.DB.First(&foundPet, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return foundPet, nil
}

// Create a new pet in the database
func (r *petRepository) Create(pet *model.Pet) (*model.Pet, error) {
	if err := r.DB.Create(&pet).Error; err != nil {
		return nil, err
	}
	return pet, nil
}

// Update a pet by ID with partial data
func (r *petRepository) Update(petId string, updateData map[string]interface{}) (*model.Pet, error) {
	var pet model.Pet
	if err := r.DB.Model(&pet).Where("id = ?", petId).Updates(updateData).Error; err != nil {
		return nil, err
	}

	// Retrieve the updated pet to return it
	if err := r.DB.First(&pet, "id = ?", petId).Error; err != nil {
		return nil, err
	}

	return &pet, nil
}