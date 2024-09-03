package repository

import (
	"backend/internal/domain/model"

	"gorm.io/gorm"
)

type IChatRepository interface {
	Create(chatData model.ChatRoom) error
	FindByUserID(userId string) (*[]model.ChatRoom, error)
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

func (r *chatRepository) FindByUserID(userId string) (*[]model.ChatRoom, error) {
	foundChatRoom := new([]model.ChatRoom)
	if err := r.DB.Find(&foundChatRoom, "user_id1 = ? OR user_id2 = ?", userId, userId).Error; err != nil {
		return nil, err
	}
	return foundChatRoom, nil
}
