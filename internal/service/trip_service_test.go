package service

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-pr-review-practice/internal/api"
)

func TestLoadTripsReturnsRecentTrips(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rides" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprint(w, `[{"id":"ride-1"}]`); err != nil {
			t.Errorf("write response: %v", err)
		}
	}))
	t.Cleanup(server.Close)
	svc := NewTripService(api.NewClient(server.Client(), server.URL))
	rides, err := svc.LoadTrips(context.Background())
	if err != nil {
		t.Fatalf("LoadTrips() error = %v", err)
	}
	if len(rides) != 1 || rides[0].ID != "ride-1" {
		t.Fatalf("LoadTrips() = %v, want ride-1", rides)
	}
}
