package service

import (
	"HyperLightLogistics-Go/internal/service/proto"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type GeocodingService struct {
	apiKey string
	client *http.Client
}

func NewGeocodingService(apiKey string) *GeocodingService {
	return &GeocodingService{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (g *GeocodingService) GetCoordinates(address *proto.DeliveryAddress) (float32, float32, error) {
	query := fmt.Sprintf("%s, %s, %s, %s, %s", address.Street, address.City, address.State, address.Zipcode, address.Country)
	encodedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("https://api.openrouteservice.org/geocode/search?api_key=%s&text=%s", g.apiKey, encodedQuery)

	resp, err := g.client.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("error fetching coordinates: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("failed to fetch coordinates, status: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, fmt.Errorf("error decoding JSON: %v", err)
	}

	if features, ok := result["features"].([]interface{}); ok && len(features) > 0 {
		if feature, ok := features[0].(map[string]interface{}); ok {
			if geometry, ok := feature["geometry"].(map[string]interface{}); ok {
				if coordinates, ok := geometry["coordinates"].([]interface{}); ok && len(coordinates) >= 2 {
					longitude := coordinates[0].(float64)
					latitude := coordinates[1].(float64)
					return float32(longitude), float32(latitude), nil
				}
			}
		}
	}

	return 0, 0, fmt.Errorf("coordinates not found in response")
}
