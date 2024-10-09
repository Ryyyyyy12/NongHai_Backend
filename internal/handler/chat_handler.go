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
		if err.Error() == "chat room already exists" {
			return c.JSON(response.InfoResponse{
				Success: true,
				Data:    "Chat room already exists",
			})
		}
		return err
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    "Chat room created",
	})
}

func (h ChatHandler) GetChatRoom(c *fiber.Ctx) error {
	body := new(dto.GetChatRoomBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	resp, err := h.chatService.GetChatRoom(*body)
	if err != nil {
		return err
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    resp,
	})
}

func (h ChatHandler) GetCurrentUserChatRoom(c *fiber.Ctx) error {
	body := new(dto.GetCurrentUserChatRoomBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	resp, err := h.chatService.GetCurrentUserChatRoom(*body)
	if err != nil {
		return err
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    resp,
	})
}

func (h ChatHandler) ReadChat(c *fiber.Ctx) error {
	body := new(dto.ReadChatRoomBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	if err := h.chatService.ReadChat(*body); err != nil {
		return err
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    "Chat read",
	})
}

func (h ChatHandler) SetUnread(c *fiber.Ctx) error {
	body := new(dto.SendMessageBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	if err := h.chatService.SetUnread(*body); err != nil {
		return err
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    "Chat unread",
	})
}
