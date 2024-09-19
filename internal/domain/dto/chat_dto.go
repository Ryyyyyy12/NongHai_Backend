package dto

type CreateChatRoomBody struct {
	ChatID  *string `json:"chat_id" validate:"required"`
	UserID1 *string `json:"user_id_1" validate:"required"`
	UserID2 *string `json:"user_id_2" validate:"required"`
}

type GetChatRoomBody struct {
	UserID *string `json:"user_id" validate:"required"`
}

type GetCurrentUserChatRoomBody struct {
	ChatID *string `json:"chat_id" validate:"required"`
}

type ReadChatRoomBody struct {
	ChatID   *string `json:"chat_id" validate:"required"`
	SenderID *string `json:"sender_id" validate:"required"`
}

type SendMessageBody struct {
	ChatID   *string `json:"chat_id" validate:"required"`
	SenderID *string `json:"sender_id" validate:"required"`
}
