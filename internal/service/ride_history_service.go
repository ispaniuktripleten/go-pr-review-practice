package service

import (
	"context"
	"fmt"

	"example.com/go-pr-review-practice/internal/model"
)

// RideRepository provides the data needed by RideHistoryService.
type RideRepository interface {
	RecentRides(context.Context) ([]model.Ride, error)
}

// RideHistoryService loads recent trips through a repository.
type RideHistoryService struct{ repository RideRepository }

// NewRideHistoryService constructs a service with its repository.
func NewRideHistoryService(repository RideRepository) *RideHistoryService {
	return &RideHistoryService{repository: repository}
}

// LoadRides returns recent rides or the repository error with context.
func (s *RideHistoryService) LoadRides(ctx context.Context) ([]model.Ride, error) {
	rides, err := s.repository.RecentRides(ctx)
	if err != nil {
		return nil, fmt.Errorf("load rides: %w", err)
	}
	return rides, nil
}
