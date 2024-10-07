package service

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/helper"
	"backend/internal/domain/model"
	"backend/internal/repository"
)

type IUserService interface {
	GetUserInfo(userId string) (*dto.UserInfoResponse, error)  // Return DTO response
	Create(user *model.User) (*model.User, error)
}

type userService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// GetUserInfo retrieves a user by ID and preloads the Pets field
func (s *userService) GetUserInfo(userId string) (*dto.UserInfoResponse, error) {
	user, err := s.userRepo.FindByIdWithPets(userId)
	if err != nil {
		return nil, err
	}

	// Convert model.User and model.Pet to corresponding DTOs
	var petsDto []dto.CreatePetBody
	for _, pet := range user.Pets {
		petDto := dto.CreatePetBody{
			ID:          pet.ID.String(),
			CreatedAt:  pet.CreatedAt,
			UpdatedAt:  pet.UpdatedAt,
			UserID:      pet.UserID,
			Name:        pet.Name,
			AnimalType:  pet.AnimalType,
			Breed:       pet.Breed,
			DateOfBirth: dto.CustomDate{Time: pet.DateOfBirth},
			Age:         helper.CalculateAge(pet.DateOfBirth), // Set the calculated Age here
			Sex:         pet.Sex,
			Weight:      pet.Weight,
			HairColor:   pet.HairColor,
			BloodType:   pet.BloodType,
			Note:        pet.Note,
			Image:       pet.Image,
		}
		petsDto = append(petsDto, petDto)
	}

	// Prepare the DTO for the user info response
	userInfo := &dto.UserInfoResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Phone:    user.Phone,
		Address:  user.Address,
		Latitude: user.Latitude,
		Longitude: user.Longitude,
		Image:    user.Image,
		Pets:     petsDto,
	}

	return userInfo, nil
}

// Create a new user in the repository
func (s *userService) Create(user *model.User) (*model.User, error) {
	newUser, err := s.userRepo.Create(*user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
