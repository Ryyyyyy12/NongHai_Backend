package handler

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/response"
	"backend/internal/service"
	"backend/internal/util/text"

	"github.com/gofiber/fiber/v2"
)

type ChatHandler struct {
	chatService service.IChatService
}

func NewChatHandler(
	chatService service.IChatService,
) ChatHandler {
	return ChatHandler{
		chatService: chatService,
	}
}

func (h ChatHandler) CreateChatRoom(c *fiber.Ctx) error {
	body := new(dto.CreateChatRoomBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	if err := h.chatService.CreateChatRoom(*body); err != nil {
		return err
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    "Chat room created",
	})
}