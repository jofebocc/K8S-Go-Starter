package main

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()
	inmem := repository.NewInMemRepository()
	svc := service.NewTripService(inmem)
	fare := &domain.RideFareModel{
		UserID:            "user123",
		PackageSlug:       "standard",
		TotalPriceInCents: 1500,
		ExpiresAt:         time.Now().Add(15 * time.Minute),
	}
	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		panic(err)
	}
	println("Created trip with fare ID:", t.UserID)

	// Keep the application running to observe behavior
	for {
		time.Sleep(1 * time.Second)
	}
}
