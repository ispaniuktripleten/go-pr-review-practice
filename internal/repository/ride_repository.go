package repository

import (
	"context"

	"example.com/go-pr-review-practice/internal/api"
	"example.com/go-pr-review-practice/internal/model"
)

// RideRepository obtains rides through the API client.
type RideRepository struct{ client *api.Client }

// NewRideRepository constructs the data-access adapter.
func NewRideRepository(client *api.Client) *RideRepository { return &RideRepository{client: client} }

// RecentRides returns the user's recent rides.
func (r *RideRepository) RecentRides(ctx context.Context) ([]model.Ride, error) {
	return r.client.RecentRides(ctx)
}
