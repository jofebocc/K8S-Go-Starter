package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InMemRepository struct {
	trips      map[string]*domain.TripModel
	ridesFares map[string]*domain.RideFareModel
}

// Delete implements domain.TripRepository.
func (r *InMemRepository) Delete(id primitive.ObjectID) error {
	panic("unimplemented")
}

// FindByID implements domain.TripRepository.
func (r *InMemRepository) FindByID(id primitive.ObjectID) (*domain.TripModel, error) {
	panic("unimplemented")
}

// Update implements domain.TripRepository.
func (r *InMemRepository) Update(trip domain.TripModel) error {
	panic("unimplemented")
}

func NewInMemRepository() *InMemRepository {
	return &InMemRepository{
		trips:      make(map[string]*domain.TripModel),
		ridesFares: make(map[string]*domain.RideFareModel),
	}
}

func (r *InMemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	// Implementation here
	r.trips[trip.ID.Hex()] = trip
	return trip, nil
}
