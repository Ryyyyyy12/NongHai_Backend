package helper

import (
	"time"

	"github.com/google/uuid"
)

type ModelBase struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;not null;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gor:"not null"`
}
