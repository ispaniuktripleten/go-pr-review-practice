package service

import (
	"context"
	"fmt"

	"example.com/go-pr-review-practice/internal/api"
	"example.com/go-pr-review-practice/internal/model"
)

// TripService loads a user's trips.
type TripService struct{ client *api.Client }

// NewTripService constructs a service with its data client.
func NewTripService(client *api.Client) *TripService { return &TripService{client: client} }

// LoadTrips retrieves the recent trips.
func (s *TripService) LoadTrips(ctx context.Context) ([]model.Ride, error) {
	rides, err := s.client.RecentRides(ctx)
	if err != nil {
		return nil, fmt.Errorf("load trips: %w", err)
	}
	return rides, nil
}
