package model

import (
	enum2 "backend/internal/domain/enum"
	"backend/internal/domain/helper"
	"time"
)

type Pet struct {
	helper.ModelBase `json:"-"`
	UserID           string           `json:"user_id" gorm:"not null;foreignKey:ID;references:ID"`
	Name             string           `json:"name" gorm:"not null"`
	AnimalType       enum2.AnimalType `json:"animal_type" gorm:"not null"`
	Breed            string           `json:"breed" gorm:"not null"`
	DateOfBirth      time.Time        `json:"date_of_birth" gorm:"not null"`
	Age         	 string    		  `json:"age" gorm:"-"` // Calculated field, not persisted in DB
	Sex              enum2.Sex        `json:"sex" gorm:"not null"`
	Weight           float64          `json:"weight" gorm:"not null"`
	HairColor        string           `json:"hair_color" gorm:"not null"`
	BloodType        string           `json:"blood_type" gorm:"not null"`
	Note             string           `json:"note"`
	Image            string           `json:"image"`
	Tracking         []Tracking       `json:"tracking" gorm:"foreignKey:PetID;references:ID"`
}
