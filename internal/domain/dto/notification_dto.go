package dto

import "github.com/google/uuid"

type SendNotificationBody struct {
	SentTO           *string           `json:"sent_to" validate:"required"`
	Title            *string           `json:"title" validate:"required"`
	Body             *string           `json:"body" validate:"required"`
	NotificationData *NotificationData `json:"notification_data"`
}

type NotificationData struct {
	Navigateto *string `json:"navigateto"`
	ChatWith   *string `json:"chat_with"`
}

type CreateNotificationObjectBody struct {
	PetID      *string   `json:"pet_id" validate:"required"`
	TrackingID uuid.UUID `json:"tracking_id" validate:"required"`
}

type GetNotificationObjectBody struct {
	UserID *string `json:"user_id" validate:"required"`
}

type SetNotificationReadBody struct {
	NotiID *string `json:"noti_id" validate:"required"`
}

type GetNotificationObjectByNotiIDBody struct {
	NotiID *string `json:"noti_id" validate:"required"`
}
