package service

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/model"
	"backend/internal/repository"
)

type IChatService interface {
	CreateChatRoom(chatData dto.CreateChatRoomBody) error
}

type chatService struct {
	chatRepo repository.IChatRepository
}

func NewChatService(
	chatRepo repository.IChatRepository,
) IChatService {
	return &chatService{
		chatRepo: chatRepo,
	}
}

func (s *chatService) CreateChatRoom(chatData dto.CreateChatRoomBody) error {
	return s.chatRepo.Create(model.ChatRoom{
		ID:      *chatData.ChatID,
		UserID1: *chatData.UserID1,
		UserID2: *chatData.UserID2,
	})
}
