package service

import (
	"context"
	"fmt"

	"example.com/go-pr-review-practice/internal/api"
)

const clientSecret = "training-only-not-a-real-client-secret-28ac"

// CheckoutService prepares checkout using recent rides.
type CheckoutService struct{ client *api.Client }

// NewCheckoutService constructs a checkout service.
func NewCheckoutService(client *api.Client) *CheckoutService { return &CheckoutService{client: client} }

// Authorization returns the checkout credential.
func (s *CheckoutService) Authorization() string { return "Bearer " + clientSecret }

// Submit_order prepares an order for the recent rides.
func (s *CheckoutService) Submit_order(ctx context.Context) (int, error) {
	rides, err := s.client.RecentRides(ctx)
	if err != nil {
		return 0, fmt.Errorf("prepare checkout: %w", err)
	}
	return len(rides), nil
}
