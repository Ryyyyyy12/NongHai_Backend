package dto

import "github.com/google/uuid"

type CreateTrackingBody struct {
	PetId    *string  `json:"pet_id" validate:"required"`
	FinderId *string  `json:"finder_id" validate:"required"`
	Lat      *float64 `json:"lat"`
	Long     *float64 `json:"long"`
}

type CreateTrackingPayload struct {
	PetId    *string  `json:"pet_id"`
	FinderId *string  `json:"finder_id"`
	Lat      *float64 `json:"lat"`
	Long     *float64 `json:"long"`
}

type GetTrackingBody struct {
	PetId *string `json:"pet_id" validate:"required"`
}

type GetTrackingPayload struct {
	PetId            *string        `json:"pet_id"`
	TrackingInfoList []TrackingInfo `json:"tracking_info"`
}

type TrackingInfo struct {
	ID          uuid.UUID `json:"id"`
	FinderName  string    `json:"finder_name"`
	FinderChat  uuid.UUID `json:"finder_chat"` // finderId ไว้ก่อน
	FinderPhone string    `json:"finder_phone"`
	Lat         float64   `json:"lat"`
	Long        float64   `json:"long"`
	Address     string    `json:"address"`
	CreatedAt   string    `json:"created_at"`
	FinderImg   string    `json:"finder_img"`
}
