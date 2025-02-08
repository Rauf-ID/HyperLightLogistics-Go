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
	"HyperLightLogistics-Go/internal/service/transport"
	"HyperLightLogistics-Go/kafka"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

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

	transportServices := []transport.TransportService{
		transport.NewDroneService(db),
		transport.NewVanService(db),
		transport.NewTruckService(db),
		transport.NewFlightService(db),
	}

	deliveryService := service.NewDeliveryService(transportServices, inventoryService, geocodingService, routeService)

	kafkaConsumer, err := kafka.NewKafkaConsumer(cfg.Kafka.Brokers, cfg.Kafka.GroupID, cfg.Kafka.Topic, deliveryService)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go kafkaConsumer.StartConsuming(ctx)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &service.DeliveryOptionsServer{
		DeliveryService: deliveryService,
	}
	proto.RegisterDeliveryOptionsServiceServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
