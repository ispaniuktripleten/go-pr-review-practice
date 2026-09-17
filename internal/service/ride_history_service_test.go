package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"example.com/go-pr-review-practice/internal/model"
)

type fakeRideRepository struct {
	rides []model.Ride
	err   error
}

func (f fakeRideRepository) RecentRides(context.Context) ([]model.Ride, error) { return f.rides, f.err }

func TestLoadRidesReturnsRepositoryRides(t *testing.T) {
	want := []model.Ride{{ID: "ride-1"}, {ID: "ride-2"}}
	svc := NewRideHistoryService(fakeRideRepository{rides: want})
	got, err := svc.LoadRides(context.Background())
	if err != nil {
		t.Fatalf("LoadRides() error = %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("LoadRides() = %v, want %v", got, want)
	}
}

func TestLoadRidesPreservesRepositoryError(t *testing.T) {
	want := errors.New("repository unavailable")
	svc := NewRideHistoryService(fakeRideRepository{err: want})
	_, err := svc.LoadRides(context.Background())
	if !errors.Is(err, want) {
		t.Fatalf("LoadRides() error = %v, want wrapped %v", err, want)
	}
}
