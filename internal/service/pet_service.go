package service

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/helper"
	"backend/internal/domain/model"
	"backend/internal/repository"
)

type IPetService interface {
	GetPetInfo(petId string) (*dto.CreatePetBody, error)
	Create(pet *model.Pet) (*model.Pet, error)
	UpdatePet(petId string, updateData map[string]interface{}) (*model.Pet, error)
}

type petService struct {
	petRepo repository.IPetRepository
}

func NewPetService(petRepo repository.IPetRepository) IPetService {
	return &petService{
		petRepo: petRepo,
	}
}

// GetPetInfo retrieves pet information by ID and calculates the age
func (s *petService) GetPetInfo(petId string) (*dto.CreatePetBody, error) {
	pet, err := s.petRepo.FindById(petId)
	if err != nil {
		return nil, err
	}

	// Convert model.Pet to dto.CreatePetBody and calculate the age
	petDto := &dto.CreatePetBody{
		ID:          pet.ID.String(),
		CreatedAt:  pet.CreatedAt,
		UpdatedAt:  pet.UpdatedAt,
		UserID:      pet.UserID,
		Name:        pet.Name,
		AnimalType:  pet.AnimalType,
		Breed:       pet.Breed,
		DateOfBirth: dto.CustomDate{Time: pet.DateOfBirth}, // wrap date as CustomDate
		Age:         helper.CalculateAge(pet.DateOfBirth),
		Sex:         pet.Sex,
		Weight:      pet.Weight,
		HairColor:   pet.HairColor,
		BloodType:   pet.BloodType,
		Eyes: pet.Eyes,
		Status: pet.Status,
		Note:        pet.Note,
		Image:       pet.Image,
	}

	return petDto, nil
}

// Create a new pet in the repository
func (s *petService) Create(pet *model.Pet) (*model.Pet, error) {
	newPet, err := s.petRepo.Create(pet)
	if err != nil {
		return nil, err
	}
	return newPet, nil
}

// UpdatePet updates pet details partially
func (s *petService) UpdatePet(petId string, updateData map[string]interface{}) (*model.Pet, error) {
	// Call the repository's Update method with the pet ID and update data
	updatedPet, err := s.petRepo.Update(petId, updateData)
	if err != nil {
		return nil, err
	}
	return updatedPet, nil
}

// func (s *petService) UpdatePet(petId string, updateBody *dto.UpdatePetBody) error {
// 	// Find the existing pet record
// 	existingPet, err := s.petRepo.FindById(petId)
// 	if err != nil {
// 		return err
// 	}

// 	// Update only the fields that are provided in the request
// 	if updateBody.Name != nil {
// 		existingPet.Name = *updateBody.Name
// 	}
// 	if updateBody.AnimalType != nil {
// 		existingPet.AnimalType = *updateBody.AnimalType
// 	}
// 	if updateBody.Breed != nil {
// 		existingPet.Breed = *updateBody.Breed
// 	}
// 	if updateBody.DateOfBirth != nil {
// 		existingPet.DateOfBirth = updateBody.DateOfBirth.Time
// 	}
// 	if updateBody.Sex != nil {
// 		existingPet.Sex = *updateBody.Sex
// 	}
// 	if updateBody.Weight != nil {
// 		existingPet.Weight = *updateBody.Weight
// 	}
// 	if updateBody.HairColor != nil {
// 		existingPet.HairColor = *updateBody.HairColor
// 	}
// 	if updateBody.BloodType != nil {
// 		existingPet.BloodType = *updateBody.BloodType
// 	}
// 	if updateBody.Eyes != nil {
// 		existingPet.Eyes = *updateBody.Eyes
// 	}
// 	if updateBody.Status != nil {
// 		existingPet.Status = *updateBody.Status
// 	}
// 	if updateBody.Note != nil {
// 		existingPet.Note = *updateBody.Note
// 	}
// 	if updateBody.Image != nil {
// 		existingPet.Image = *updateBody.Image
// 	}

// 	// Save the updated pet record
// 	return s.petRepo.Update(existingPet)
// }
