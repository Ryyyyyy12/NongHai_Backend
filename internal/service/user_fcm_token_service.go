package service

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/model"
	"backend/internal/repository"
	"errors"
	"fmt"
)

type IUserFCMTokenService interface {
	CreateUserFCMToken(body dto.UserFCMTokenBody) error
	GetUserFCMToken(userID string) ([]string, error)
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
	fmt.Print(body.Token)
	token, err := s.TokenRepo.GetTokenByToken(body.Token)
	if err != nil && err.Error() != "record not found" {
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

func (s *userFCMTokenService) GetUserFCMToken(userID string) ([]string, error) {
	token, err := s.TokenRepo.GetTokenByUserID(userID)
	if err != nil {
		return nil, err
	}
	var tokens []string
	for _, t := range *token {
		tokens = append(tokens, t.Token)
	}

	fmt.Print(tokens)
	return tokens, nil
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
