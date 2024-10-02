package repository

import (
	"backend/internal/domain/model"

	"gorm.io/gorm"
)

type INotificationRepository interface {
	CreateNotificationObject(notificationObject model.Notification) error
	GetNotificationObjectByUserID(userID string) ([]*model.Notification, error)
	GetNotificationObjectByNotiID(notificationID string) (*model.Notification, error)
	UpdateNotificationObject(notificationObject model.Notification) error
}

type notificationRepository struct {
	DB gorm.DB
}

func NewNotificationRepository(db gorm.DB) INotificationRepository {
	return &notificationRepository{
		DB: db,
	}
}

func (r *notificationRepository) CreateNotificationObject(notificationObject model.Notification) error {
	return r.DB.Create(&notificationObject).Error
}

func (r *notificationRepository) GetNotificationObjectByUserID(userID string) ([]*model.Notification, error) {
	notiObject := new([]*model.Notification)
	err := r.DB.Find(&notiObject, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return *notiObject, nil
}

func (r *notificationRepository) GetNotificationObjectByNotiID(notificationID string) (*model.Notification, error) {
	notiObject := new(model.Notification)
	err := r.DB.Find(&notiObject, "id = ?", notificationID).Error
	if err != nil {
		return nil, err
	}
	return notiObject, nil
}

func (r *notificationRepository) UpdateNotificationObject(notificationObject model.Notification) error {
	return r.DB.Save(&notificationObject).Error

}
