package handler

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/response"
	"backend/internal/service"
	"backend/internal/util/text"
	"backend/loaders/config"
	"github.com/gofiber/fiber/v2"
	"github.com/kelvins/geocoder"
)

type TrackingHandler struct {
	trackingService service.ITrackingService
	userService     service.IUserService
}

func NewTrackingHandler(trackingService service.ITrackingService, userService service.IUserService) TrackingHandler {
	return TrackingHandler{
		trackingService: trackingService,
		userService:     userService,
	}
}

func (h TrackingHandler) CreateTracking(c *fiber.Ctx) error {
	body := new(dto.CreateTrackingBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	resp, err := h.trackingService.Create(*body.PetId, *body.FinderId, *body.Lat, *body.Long)
	if err != nil {
		return err
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    resp,
	})
}

func (h TrackingHandler) GetTracking(c *fiber.Ctx) error {
	geocoder.ApiKey = config.Conf.GoogleAPIKey

	body := new(dto.GetTrackingBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	//Get all tracking
	trackings, err := h.trackingService.GetAllById(*body.PetId)
	if err != nil {
		return err
	}

	trackingInfoList := make([]dto.TrackingInfo, len(*trackings))

	for i, tracking := range *trackings {

		// Get finder info
		user, err := h.userService.GetUserInfo(tracking.FinderID)
		if err != nil {
			return err
		}

		// Generate address from lat , long
		location := geocoder.Location{
			Latitude:  tracking.Latitude,
			Longitude: tracking.Longitude,
		}

		// request geocoding api
		var respAddress string
		address, err := geocoder.GeocodingReverse(location)
		if err != nil {
			respAddress = "No results found"
		} else {
			respAddress = address[0].FormattedAddress
		}

		// map tracking info
		trackingInfo := dto.TrackingInfo{
			ID:          tracking.ID,
			FinderName:  user.Username,
			FinderChat:  user.ID,
			FinderPhone: user.Phone,
			Lat:         tracking.Latitude,
			Long:        tracking.Longitude,
			Address:     respAddress,
			CreatedAt:   tracking.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		trackingInfoList[i] = trackingInfo
	}

	// map response
	resp := dto.GetTrackingPayload{
		PetId:            body.PetId,
		TrackingInfoList: trackingInfoList,
	}

	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    resp,
	})
}
