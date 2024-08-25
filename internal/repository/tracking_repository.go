package repository

import (
	"backend/internal/domain/model"
	"gorm.io/gorm"
)

type trackingRepository struct {
	DB gorm.DB
}

type ITrackingRepository interface {
	Create(tracking model.Tracking) (newTracking *model.Tracking, err error)
	FindByPetId(PetId string) (tracking *[]model.Tracking, err error)
}

func NewTrackingRepository(db gorm.DB) ITrackingRepository {
	return &trackingRepository{
		DB: db,
	}
}

func (r *trackingRepository) Create(tracking model.Tracking) (newTracking *model.Tracking, err error) {
	if err = r.DB.Create(&tracking).Error; err != nil {
		return nil, err
	}
	return &tracking, nil
}

func (r *trackingRepository) FindByPetId(PetId string) (tracking *[]model.Tracking, err error) {
	foundTracking := new([]model.Tracking)
	if err = r.DB.Find(&foundTracking, "pet_id = ?", PetId).Error; err != nil {
		return nil, err
	}
	return foundTracking, nil
}
