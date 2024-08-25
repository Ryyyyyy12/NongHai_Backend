package service

import (
	"backend/internal/domain/model"
	"backend/internal/repository"
)

type IUserService interface {
	GetUserInfo(userId string) (user *model.User, err error)
}

type userService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetUserInfo(userId string) (user *model.User, err error) {
	return s.userRepo.FindById(userId)
}
