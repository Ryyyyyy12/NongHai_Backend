package service

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/model"
	"backend/internal/repository"
)

type IChatService interface {
	CreateChatRoom(chatData dto.CreateChatRoomBody) error
	GetChatRoom(chatData dto.GetChatRoomBody) (*[]model.ChatRoom, error)
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

func (s *chatService) GetChatRoom(chatData dto.GetChatRoomBody) (*[]model.ChatRoom, error) {
	chatRoom, err := s.chatRepo.FindByUserID(*chatData.UserID)
	if err != nil {
		return nil, err
	}

	return chatRoom, nil
}
