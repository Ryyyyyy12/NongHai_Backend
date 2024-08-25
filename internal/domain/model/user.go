package model

import (
	"backend/internal/domain/helper"
)

type User struct {
	helper.ModelBase `json:"-"`
	Username         string         `json:"username" gorm:"not null"`
	Name             string         `json:"name" gorm:"not null"`
	Surname          string         `json:"surname" gorm:"not null"`
	Email            string         `json:"email" gorm:"not null"`
	Phone            string         `json:"phone" gorm:"not null"`
	Address          string         `json:"address" gorm:"not null"`
	Latitude         float64        `json:"latitude"`
	Longitude        float64        `json:"longitude"`
	Image            string         `json:"image"`
	Pets             []Pet          `json:"pets" gorm:"foreignKey:UserID;references:ID"`
	Notifications    []Notification `json:"notifications" gorm:"foreignKey:UserID;references:ID"`
	Tokens           []Token        `json:"tokens" gorm:"foreignKey:UserID;references:ID"`
	ChatRooms        []ChatRoom     `json:"chat_rooms" gorm:"foreignKey:UserID1;references:ID"`
}
