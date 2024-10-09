package service

import (
	"backend/internal/domain/model"
	"backend/internal/repository"
	"sort"
	"strings"
)

type ITrackingService interface {
	Create(petId string, finderId string, lat float64, long float64) (tracking *model.Tracking, err error)
	GetAllById(petId string) (tracking *[]model.Tracking, err error)
}

type trackingService struct {
	trackingRepo repository.ITrackingRepository
	userRepo     repository.IUserRepository
	petRepo      repository.IPetRepository
}

func NewTrackingService(trackingRepo repository.ITrackingRepository, userRepo repository.IUserRepository, petRepo repository.IPetRepository) ITrackingService {
	return &trackingService{
		trackingRepo: trackingRepo,
		userRepo:     userRepo,
		petRepo:      petRepo,
	}
}

func (s *trackingService) Create(petId string, finderId string, lat float64, long float64) (tracking *model.Tracking, err error) {

	if _, err := s.userRepo.FindById(finderId); err != nil {
		return nil, err
	}

	return s.trackingRepo.Create(model.Tracking{
		PetID:     strings.TrimSpace(petId),
		FinderID:  strings.TrimSpace(finderId),
		Latitude:  lat,
		Longitude: long,
	})
}

func (s *trackingService) GetAllById(petId string) (tracking *[]model.Tracking, err error) {
	if _, err := s.petRepo.FindById(petId); err != nil {
		return nil, err
	}
	foundTracking, err := s.trackingRepo.FindByPetId(petId)
	if err != nil {
		return nil, err
	}
	//order by created_at
	sort.Slice(*foundTracking, func(i, j int) bool {
		return (*foundTracking)[i].CreatedAt.After((*foundTracking)[j].CreatedAt)
	})
	return foundTracking, nil
}
