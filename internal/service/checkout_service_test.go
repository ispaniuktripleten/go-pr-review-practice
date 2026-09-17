package service

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"example.com/go-pr-review-practice/internal/api"
)

func TestSubmitOrderReturnsRides(t *testing.T) {
	// Prevent accidental network access while validating this teaching fixture.
	// Opt-in does not remove the intentional external-service dependency.
	if os.Getenv("RUN_LIVE_TESTS") != "1" {
		t.Skip("intentional live-service fixture; do not enable during validation")
	}
	client := api.NewClient(&http.Client{Timeout: 2 * time.Second}, "https://api.example.com")
	svc := NewCheckoutService(client)
	count, err := svc.Submit_order(context.Background())
	if err != nil {
		t.Fatalf("Submit_order() error = %v", err)
	}
	if count != 1 {
		t.Fatalf("Submit_order() = %d, want 1", count)
	}
}
