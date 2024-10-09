package service

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/model"
	"backend/internal/repository"
	"errors"
)

type IChatService interface {
	CreateChatRoom(chatData dto.CreateChatRoomBody) error
	GetChatRoom(chatData dto.GetChatRoomBody) (*[]model.ChatRoom, error)
	GetCurrentUserChatRoom(chatData dto.GetCurrentUserChatRoomBody) (*model.ChatRoom, error)
	ReadChat(chatData dto.ReadChatRoomBody) error
	SetUnread(chatData dto.SendMessageBody) error
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
	_, err := s.chatRepo.FindByChatID(*chatData.ChatID)
	if err == nil {
		return errors.New("chat room already exists")
	}

	err = s.chatRepo.Create(model.ChatRoom{
		ID:      *chatData.ChatID,
		UserID1: *chatData.UserID1,
		UserID2: *chatData.UserID2,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *chatService) GetChatRoom(chatData dto.GetChatRoomBody) (*[]model.ChatRoom, error) {
	chatRoom, err := s.chatRepo.FindByUserID(*chatData.UserID)
	if err != nil {
		return nil, err
	}

	return chatRoom, nil
}

func (s *chatService) GetCurrentUserChatRoom(chatData dto.GetCurrentUserChatRoomBody) (*model.ChatRoom, error) {
	chatRoom, err := s.chatRepo.FindByChatID(*chatData.ChatID)
	if err != nil {
		return nil, err
	}

	return chatRoom, nil
}

func (s *chatService) ReadChat(chatData dto.ReadChatRoomBody) error {
	chatRoom, err := s.chatRepo.FindByChatID(*chatData.ChatID)
	if err != nil {
		return err
	}

	if *chatData.SenderID != chatRoom.UserID1 && *chatData.SenderID != chatRoom.UserID2 {
		return errors.New("user not in chat room")
	}

	if *chatData.SenderID == chatRoom.UserID1 {
		// Set user 1 to read
		chatRoom.IsUser1Read = true
	}

	if *chatData.SenderID == chatRoom.UserID2 {
		//Set user 2 to read
		chatRoom.IsUser2Read = true
	}

	return s.chatRepo.UpdateChat(*chatRoom)
}

func (s *chatService) SetUnread(chatData dto.SendMessageBody) error {
	chatRoom, err := s.chatRepo.FindByChatID(*chatData.ChatID)
	if err != nil {
		return err
	}

	if *chatData.SenderID != chatRoom.UserID1 && *chatData.SenderID != chatRoom.UserID2 {
		return errors.New("user not in chat room")
	}

	if *chatData.SenderID == chatRoom.UserID1 {
		// Set the other user to unread
		chatRoom.IsUser2Read = false
	}

	if *chatData.SenderID == chatRoom.UserID2 {
		// Set the other user to unread
		chatRoom.IsUser1Read = false
	}

	return s.chatRepo.UpdateChat(*chatRoom)
}
