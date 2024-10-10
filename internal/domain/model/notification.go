package model

import (
	"backend/internal/domain/helper"

	"github.com/google/uuid"
)

type Notification struct {
	helper.ModelBase `json:",inline"`
	UserID           string    `json:"user_id" gorm:"not null;foreignKey:ID;references:ID"`
	PetID            string    `json:"pet_id" gorm:"not null;foreignKey:ID;references:ID"`
	TrackingID       uuid.UUID `json:"tracking_id" gorm:"not null;foreignKey:ID;references:ID"`
	IsRead           bool      `json:"is_read" gorm:"not null"`
}
