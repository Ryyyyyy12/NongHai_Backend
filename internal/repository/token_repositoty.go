package repository

import (
	"backend/internal/domain/model"

	"gorm.io/gorm"
)

type ITokenRepository interface {
	CreateToken(tokenData model.Token) error
	GetTokenByUserID(userID string) (*[]model.Token, error)
	GetTokenByToken(token string) (*model.Token, error)
	RemoveToken(tokenData model.Token) error
}

type tokenRepository struct {
	DB gorm.DB
}

func NewTokenRepository(db gorm.DB) ITokenRepository {
	return &tokenRepository{
		DB: db,
	}
}

func (r *tokenRepository) CreateToken(tokenData model.Token) error {
	return r.DB.Create(&tokenData).Error
}

func (r *tokenRepository) GetTokenByUserID(userID string) (*[]model.Token, error) {
	foundToken := new([]model.Token)
	if err := r.DB.Find(&foundToken, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return foundToken, nil
}

func (r *tokenRepository) GetTokenByToken(token string) (*model.Token, error) {
	foundToken := new(model.Token)
	if err := r.DB.Find(&foundToken, "token = ?", token).Error; err != nil {
		return nil, err
	}
	return foundToken, nil
}

func (r *tokenRepository) RemoveToken(tokenData model.Token) error {
	return r.DB.Delete(&tokenData).Error
}
