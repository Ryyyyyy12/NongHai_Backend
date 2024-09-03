package dto

type CreateChatRoomBody struct {
	ChatID  *string `json:"chat_id" validate:"required"`
	UserID1 *string `json:"user_id_1" validate:"required"`
	UserID2 *string `json:"user_id_2" validate:"required"`
}