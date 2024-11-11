/*
 * This file is part of HyperLightLogistics-Go.
 *
 * HyperLightLogistics-Go is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * HyperLightLogistics-Go is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with HyperLightLogistics-Go.  If not, see <https://www.gnu.org/licenses/>.
 *
 * Copyright (C) 2024 Rauf Agaguliev
 */

package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
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

func (r *RouteService) CalculateDistance(clientLon, clientLat float32, warehouses []WarehouseInfo) (*WarehouseInfo, float64, error) {
	if len(warehouses) == 0 {
		return nil, 0, errors.New("no warehouses provided")
	}

	sort.Slice(warehouses, func(i, j int) bool {
		return calculateEuclideanDistance(clientLon, clientLat, warehouses[i].Longitude, warehouses[i].Latitude) <
			calculateEuclideanDistance(clientLon, clientLat, warehouses[j].Longitude, warehouses[j].Latitude)
	})

	var closestWarehouse *WarehouseInfo
	minDistance := math.MaxFloat64

	for _, warehouse := range warehouses {
		routeDistance, err := r.getDistance(clientLon, clientLat, warehouse.Longitude, warehouse.Latitude)
		if err != nil {
			log.Printf("Error fetching distance for warehouse %d: %v\n", warehouse.WarehouseID, err)
			continue
		}

		if routeDistance < minDistance {
			minDistance = routeDistance
			closestWarehouse = &warehouse
		} else {
			break
		}
	}

	if closestWarehouse == nil {
		return nil, 0, errors.New("no accessible warehouses found")
	}

	log.Println(closestWarehouse)
	log.Println(minDistance)

	return closestWarehouse, minDistance, nil
}

func (r *RouteService) getDistance(clientLon, clientLat, warehouseLon, warehouseLat float32) (float64, error) {
	url := fmt.Sprintf(
		"https://api.openrouteservice.org/v2/directions/driving-car?api_key=%s&start=%f,%f&end=%f,%f",
		r.apiKey, warehouseLon, warehouseLat, clientLon, clientLat,
	)

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
		return 0, fmt.Errorf("error decoding JSON: %v", err)
	}

	if features, ok := result["features"].([]interface{}); ok && len(features) > 0 {
		if feature, ok := features[0].(map[string]interface{}); ok {
			if properties, ok := feature["properties"].(map[string]interface{}); ok {
				if segments, ok := properties["segments"].([]interface{}); ok && len(segments) > 0 {
					if segment, ok := segments[0].(map[string]interface{}); ok {
						if distance, ok := segment["distance"].(float64); ok {
							log.Println(distance)
							return distance, nil
						}
					}
				}
			}
		}
	}

	return 0, fmt.Errorf("distance data not found in response")
}

func calculateEuclideanDistance(clientLon, clientLat, warehouseLon, warehouseLat float32) float64 {
	radius := 6371e3
	lat1Rad := float64(clientLat) * math.Pi / 180
	lat2Rad := float64(warehouseLat) * math.Pi / 180
	deltaLat := (float64(warehouseLat) - float64(clientLat)) * math.Pi / 180
	deltaLon := (float64(warehouseLon) - float64(clientLon)) * math.Pi / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return radius * c
}
