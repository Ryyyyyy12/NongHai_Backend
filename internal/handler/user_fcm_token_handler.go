package handler

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/response"
	"backend/internal/service"
	"backend/internal/util/text"

	"github.com/gofiber/fiber/v2"
)

type UserTokenHandler struct {
	userFCMTokenService service.IUserFCMTokenService
}

func NewUserTokenHandler(
	userFCMTokenService service.IUserFCMTokenService,
) UserTokenHandler {
	return UserTokenHandler{
		userFCMTokenService: userFCMTokenService,
	}
}

func (h UserTokenHandler) CreateUserFCMToken(c *fiber.Ctx) error {
	body := new(dto.UserFCMTokenBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	if err := h.userFCMTokenService.CreateUserFCMToken(*body); err != nil {
		return err
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    "Token created",
	})
}

func (h UserTokenHandler) CheckIfTokenExist(c *fiber.Ctx) error {
	body := new(dto.UserFCMTokenBody)

	tokens, err := h.userFCMTokenService.GetUserFCMToken(body.UserID)
	if err != nil {
		return err
	}

	if len(*tokens) == 0 {
		return c.JSON(response.InfoResponse{
			Success: true,
			Data:    false, // Token not exist
		})
	} else {
		return c.JSON(response.InfoResponse{
			Success: true,
			Data:    true, // Token exist
		})
	}

}

func (h UserTokenHandler) RemoveUserFCMToken(c *fiber.Ctx) error {
	body := new(dto.UserFCMTokenBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	if err := h.userFCMTokenService.RemoveUserFCMToken(*body); err != nil {
		return err
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    "Token removed",
	})
}
