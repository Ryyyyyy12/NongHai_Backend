package model

import (
	"backend/internal/domain/helper"
)

type Notification struct {
	helper.ModelBase `json:"-"`
	UserID           string `json:"user_id" gorm:"not null;foreignKey:ID;references:ID"`
	PetID            string `json:"pet_id" gorm:"not null;foreignKey:ID;references:ID"`
	TrackingID       string `json:"tracking_id" gorm:"not null;foreignKey:ID;references:ID"`
	IsRead           bool   `json:"is_read" gorm:"not null"`
}
