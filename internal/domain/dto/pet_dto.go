package dto

import (
	enum2 "backend/internal/domain/enum"
	// "backend/internal/domain/helper"
	"strings"
	"time"
)

type CreatePetBody struct {
	// helper.ModelBase
	ID          string           `json:"id"`
	CreatedAt     time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time      `json:"updated_at" gor:"not null"`
	UserID      string           `json:"user_id" gorm:"not null;foreignKey:ID;references:ID"`
	Name        string           `json:"name" gorm:"not null"`
	AnimalType  enum2.AnimalType `json:"animal_type" gorm:"not null"`
	Breed       string           `json:"breed" gorm:"not null"`
	DateOfBirth CustomDate       `json:"date_of_birth" gorm:"not null"`	
	Age         string           `json:"age,omitempty"`
	Sex         enum2.Sex        `json:"sex" gorm:"not null"`
	Weight      float64          `json:"weight" gorm:"not null"`
	HairColor   string           `json:"hair_color" gorm:"not null"`
	BloodType   string           `json:"blood_type" gorm:"not null"`
	Note        string           `json:"note"`
	Image       string           `json:"image"`
}

type CustomDate struct {
    time.Time
}

// UnmarshalJSON handles the custom unmarshaling for CustomDate
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
    str := strings.Trim(string(b), `"`)
    t, err := time.Parse("2006-01-02", str) // Accepts date in "YYYY-MM-DD" format
    if err != nil {
        return err
    }
    cd.Time = t
    return nil
}
