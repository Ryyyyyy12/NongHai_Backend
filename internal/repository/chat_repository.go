package repository

import (
	"backend/internal/domain/model"

	"gorm.io/gorm"
)

type IChatRepository interface {
	Create(chatData model.ChatRoom) error
}

type chatRepository struct {
	DB gorm.DB
}

func NewChatRepository(db gorm.DB) IChatRepository {
	return &chatRepository{
		DB: db,
	}
}

func (r *chatRepository) Create(chatData model.ChatRoom) error {
	return r.DB.Create(&chatData).Error
}
