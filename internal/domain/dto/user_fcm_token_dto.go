package dto

type UserFCMTokenBody struct {
	UserID string `json:"user_id" validate:"required"`
	Token  string `json:"token" validate:"required"`
}

type GetUserFCMTokenBody struct {
	UserID string `json:"user_id" validate:"required"`
}
