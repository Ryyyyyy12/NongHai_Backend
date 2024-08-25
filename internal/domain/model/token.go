package model

import (
	"backend/internal/domain/helper"
)

type Token struct {
	helper.ModelBase `json:"-"`
	UserID           string `json:"user_id" gorm:"not null;foreignKey:ID;references:ID"`
	Token            string `json:"token" gorm:"not null"`
}
