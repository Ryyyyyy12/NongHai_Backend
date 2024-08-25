package model

import (
	"backend/internal/domain/enum"
	"backend/internal/domain/helper"
)

type Message struct {
	helper.ModelBase `json:"-"`
	ChatRoomID       string           `json:"chat_room_id" gorm:"not null;foreignKey:ID;references:ID"`
	SenderID         string           `json:"sender_id" gorm:"not null;foreignKey:ID;references:ID"`
	MessageType      enum.MessageType `json:"message_type" gorm:"not null"`
	Message          string           `json:"message" gorm:"not null"`
	IsRead           bool             `json:"is_read" gorm:"not null"`
}
