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

package main

import (
	"HyperLightLogistics-Go/internal/config"
	"HyperLightLogistics-Go/internal/db"
	"HyperLightLogistics-Go/internal/service"
	proto "HyperLightLogistics-Go/internal/service/proto"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type DeliveryOptionsServer struct {
	proto.UnimplementedDeliveryOptionsServiceServer
	InventoryService *service.InventoryService
	RouteService     *service.RouteService
	GeocodingService *service.GeocodingService
}

func (s DeliveryOptionsServer) CalculateDeliveryOptions(ctx context.Context, req *proto.DeliveryRequest) (*proto.DeliveryResponse, error) {

	clientLon, clientLat, err := s.GeocodingService.GetCoordinates(req.DeliveryAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get coordinates: %v", err)
	}

	log.Println(clientLat, clientLon)

	for _, product := range req.Products {
		productId := product.ProductId

		warehouses, err := s.InventoryService.GetWarehousesForProduct(productId)
		if err != nil {
			return nil, err
		}

		for _, warehouse := range warehouses {
			log.Printf("For product: %d warehouse was found: %d with quantity: %d warehouse location: %s latitude: %f longitude: %f",
				warehouse.ProductID, warehouse.WarehouseID, warehouse.Quantity, warehouse.Location, warehouse.Latitude, warehouse.Longitude)
		}

		closestWarehouse, distance, err := s.RouteService.CalculateDistance(clientLon, clientLat, warehouses)
		if err != nil {
			return nil, err
		}

		_ = closestWarehouse
		_ = distance

	}

	deliveryOptions := calculateDeliveryRoutes()

	return &proto.DeliveryResponse{
		DeliveryOptions: deliveryOptions,
	}, nil
}

func calculateDeliveryRoutes() []*proto.DeliveryOptions {
	var deliveryOptions []*proto.DeliveryOptions

	deliveryOptions = append(deliveryOptions, &proto.DeliveryOptions{
		Type:         "Standard",
		DeliveryTime: "3-5 days",
		Price:        5.99,
	})

	deliveryOptions = append(deliveryOptions, &proto.DeliveryOptions{
		Type:         "Expedited",
		DeliveryTime: "1-3 days",
		Price:        15.99,
	})

	return deliveryOptions
}

func main() {
	cfg, err := config.LoadConfig("../config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	db, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.Close()

	inventoryService := service.NewInventoryService(db)
	routeService := service.NewRouteService(cfg.OpenRouteService.OpenRouteServiceAPIKey)
	geocodingService := service.NewGeocodingService(cfg.OpenRouteService.OpenRouteServiceAPIKey)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &DeliveryOptionsServer{
		InventoryService: inventoryService,
		RouteService:     routeService,
		GeocodingService: geocodingService,
	}
	proto.RegisterDeliveryOptionsServiceServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
