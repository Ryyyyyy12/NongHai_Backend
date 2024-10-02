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
