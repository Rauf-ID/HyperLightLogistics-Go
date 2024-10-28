package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type RouteService struct {
	apiKey string
	client *http.Client
}

func NewRouteService(apiKey string) *RouteService {
	return &RouteService{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (r *RouteService) CalculateDistance(clientLon, clientLat float32, warehouses []InventoryItem) (*InventoryItem, float64, error) {
	var closestWarehouse *InventoryItem
	var minDistance float64 = -1

	for _, warehouse := range warehouses {
		distance, err := r.getDistance(clientLon, clientLat, warehouse.Longitude, warehouse.Latitude)
		if err != nil {
			return nil, 0, err
		}

		if minDistance == -1 || distance < minDistance {
			minDistance = distance
			closestWarehouse = &warehouse
		}
	}

	if closestWarehouse == nil {
		return nil, 0, errors.New("no warehouses found")
	}

	return closestWarehouse, minDistance, nil
}

func (r *RouteService) getDistance(clientLon, clientLat, warehouseLon, warehouseLat float32) (float64, error) {
	url := fmt.Sprintf(
		"https://api.openrouteservice.org/v2/directions/driving-car?api_key=%s&start=%f,%f&end=%f,%f",
		r.apiKey, warehouseLon, warehouseLat, clientLon, clientLat,
	)

	// warehouse 8.681495, 49.41461,     client 8.687872, 49.420318

	resp, err := r.client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch route data: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	log.Println(resp.StatusCode)
	log.Println(result)

	return 0, nil
}
