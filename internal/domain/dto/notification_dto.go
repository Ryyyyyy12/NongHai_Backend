package dto

type SendNotificationBody struct {
	UserID string `json:"user_id" validate:"required"`
}
