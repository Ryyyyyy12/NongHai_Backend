package service

import (
	"backend/internal/domain/model"
	"backend/internal/repository"
	"fmt"
	"time"
)

type IUserService interface {
	GetUserInfo(userId string) (*model.User, error)
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

// Helper function to calculate age
func calculateAge(birthDate time.Time) string {
	now := time.Now()
	years := now.Year() - birthDate.Year()

	// If the pet's birthday hasn't occurred yet this year, subtract a year from the age
	if now.YearDay() < birthDate.YearDay() {
		years--
	}

	if years > 0 {
		if years == 1 {
			return "1 year"
		}
		return fmt.Sprintf("%d years", years)
	}
	
	// Otherwise, return age in months
	months := int(now.Sub(birthDate).Hours() / (24 * 30))
	if months == 0 || months == 1 {
		return fmt.Sprintf("%d month", months)
	}
	return fmt.Sprintf("%d months", months)
	
}

// GetUserInfo retrieves a user by ID and preloads the Pets field
func (s *userService) GetUserInfo(userId string) (*model.User, error) {
	// Fetch user from the repository, which should include preloading the pets
	user, err := s.userRepo.FindByIdWithPets(userId)
	if err != nil {
		return nil, err
	}

	// Calculate the age for each pet and store it in a custom field
	for i := range user.Pets {
		user.Pets[i].Age = calculateAge(user.Pets[i].DateOfBirth)
	}

	return user, nil
}

// Create a new user in the repository
func (s *userService) Create(user *model.User) (*model.User, error) {
	newUser, err := s.userRepo.Create(*user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
