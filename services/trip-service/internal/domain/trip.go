package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   string
	Status   string
	RideFare *RideFareModel
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
	FindByID(id primitive.ObjectID) (*TripModel, error)
	Update(trip TripModel) error
	Delete(id primitive.ObjectID) error
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*RideFareModel, error)
	FindByID(id primitive.ObjectID) (*RideFareModel, error)
	Update(trip RideFareModel) error
	Delete(id primitive.ObjectID) error
}
