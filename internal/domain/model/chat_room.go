package model

import (
	"time"
)

type ChatRoom struct {
	ID        string    `json:"id" gorm:"primary_key;not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gor:"not null"`
	UserID1   string    `json:"user_id_1" gorm:"not null;foreignKey:ID;references:ID"`
	UserID2   string    `json:"user_id_2" gorm:"not null;foreignKey:ID;references:ID"`
}
