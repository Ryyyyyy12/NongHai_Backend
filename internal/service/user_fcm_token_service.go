package service

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/model"
	"backend/internal/repository"
	"errors"
)

type IUserFCMTokenService interface {
	CreateUserFCMToken(body dto.UserFCMTokenBody) error
	GetUserFCMToken(userID string) (*[]model.Token, error)
	RemoveUserFCMToken(body dto.UserFCMTokenBody) error
}

type userFCMTokenService struct {
	TokenRepo repository.ITokenRepository
}

func NewUserFCMTokenService(
	tokenRepo repository.ITokenRepository,
) IUserFCMTokenService {
	return &userFCMTokenService{
		TokenRepo: tokenRepo,
	}
}

func (s *userFCMTokenService) CreateUserFCMToken(body dto.UserFCMTokenBody) error {
	token, err := s.TokenRepo.GetTokenByToken(body.Token)
	if err != nil {
		return err
	}

	if token != nil {
		return errors.New("token already exist")
	}

	return s.TokenRepo.CreateToken(model.Token{
		UserID: body.UserID,
		Token:  body.Token,
	})
}

func (s *userFCMTokenService) GetUserFCMToken(userID string) (*[]model.Token, error) {
	token, err := s.TokenRepo.GetTokenByUserID(userID)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *userFCMTokenService) RemoveUserFCMToken(body dto.UserFCMTokenBody) error {
	token, err := s.TokenRepo.GetTokenByToken(body.Token)
	if err != nil {
		return err
	}

	if token.UserID != body.UserID {
		return errors.New("wrong user id for token: " + body.Token)
	}

	return s.TokenRepo.RemoveToken(*token)
}
