package model

import (
	"backend/internal/domain/helper"
)

type ChatRoom struct {
	helper.ModelBase `json:"-"`
	UserID1          string    `json:"user_id_1" gorm:"not null;foreignKey:ID;references:ID"`
	UserID2          string    `json:"user_id_2" gorm:"not null;foreignKey:ID;references:ID"`
	Messages         []Message `json:"messages" gorm:"foreignKey:ChatRoomID;references:ID"`
}
