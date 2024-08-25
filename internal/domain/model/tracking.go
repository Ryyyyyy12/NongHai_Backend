package model

import (
	"backend/internal/domain/helper"
)

type Tracking struct {
	helper.ModelBase `json:"-"`
	PetID            string  `json:"pet_id" gorm:"not null;foreignKey:ID;references:ID"`
	FinderID         string  `json:"finder_id" gorm:"not null;foreignKey:ID;references:ID"`
	Latitude         float64 `json:"latitude" gorm:"not null"`
	Longitude        float64 `json:"longitude" gorm:"not null"`
}
