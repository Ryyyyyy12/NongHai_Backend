package repository

import (
	"backend/internal/domain/model"

	"gorm.io/gorm"
)

type ITokenRepository interface {
	CreateToken(tokenData model.Token) error
	GetTokenByUserID(userID string) (*[]model.Token, error)
	GetTokenByTokenAndUserID(token, user_id string) (*model.Token, error)
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

func (r *tokenRepository) GetTokenByTokenAndUserID(token, userID string) (*model.Token, error) {
	foundToken := new(model.Token)
	// Use both token and user_id in the query condition
	if err := r.DB.First(&foundToken, "token = ? AND user_id = ?", token, userID).Error; err != nil {
		return nil, err
	}

	return foundToken, nil
}


func (r *tokenRepository) RemoveToken(tokenData model.Token) error {
	return r.DB.Delete(&tokenData).Error
}
