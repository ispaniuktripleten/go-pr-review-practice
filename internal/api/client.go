package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/go-pr-review-practice/internal/model"
)

// Client retrieves rides from an upstream API.
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient constructs a client with dependencies supplied by the caller.
func NewClient(httpClient *http.Client, baseURL string) *Client {
	return &Client{httpClient: httpClient, baseURL: baseURL}
}

// RecentRides retrieves and decodes recent rides.
func (c *Client) RecentRides(ctx context.Context) ([]model.Ride, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/rides", nil)
	if err != nil {
		return nil, fmt.Errorf("create rides request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch rides: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("rides status: %d", resp.StatusCode)
	}
	var rides []model.Ride
	if err := json.NewDecoder(resp.Body).Decode(&rides); err != nil {
		return nil, fmt.Errorf("decode rides: %w", err)
	}
	return rides, nil
}
