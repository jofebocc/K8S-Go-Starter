package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

type serviceTrip struct {
	repo domain.TripRepository
}

func NewTripService(repo domain.TripRepository) *serviceTrip {
	return &serviceTrip{
		repo: repo,
	}
}

func (s *serviceTrip) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.RideFareModel, error) {
	trip := &domain.TripModel{
		UserID:   fare.UserID,
		Status:   "created",
		RideFare: fare,
	}
	createdTrip, err := s.repo.CreateTrip(ctx, trip)
	if err != nil {
		return nil, err
	}
	return createdTrip.RideFare, nil
}
