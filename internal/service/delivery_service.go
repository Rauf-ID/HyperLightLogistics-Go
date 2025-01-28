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
	"HyperLightLogistics-Go/internal/service/proto"
	"HyperLightLogistics-Go/internal/service/transport"
	"context"
	"fmt"
	"log"
)

type DeliveryOptionsServer struct {
	proto.UnimplementedDeliveryOptionsServiceServer
	DeliveryService *DeliveryService
}

type DeliveryService struct {
	TransportServices []transport.TransportService
	InventoryService  *InventoryService
	GeocodingService  *GeocodingService
	RouteService      *RouteService
}

func NewDeliveryService(
	transportServices []transport.TransportService,
	inventoryService *InventoryService,
	geocodingService *GeocodingService,
	routeService *RouteService,
) *DeliveryService {
	return &DeliveryService{
		TransportServices: transportServices,
		InventoryService:  inventoryService,
		GeocodingService:  geocodingService,
		RouteService:      routeService,
	}
}

func (s *DeliveryOptionsServer) CalculateDeliveryOptions(ctx context.Context, req *proto.DeliveryRequest) (*proto.DeliveryResponse, error) {
	var productDeliveryOptions []*proto.ProductDeliveryOptions

	clientLon, clientLat, err := s.DeliveryService.GeocodingService.GetCoordinates(req.DeliveryAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get coordinates: %v", err)
	}

	for _, product := range req.Products {
		productId := product.ProductId

		warehouses, err := s.DeliveryService.InventoryService.GetWarehousesInfoByProduct(productId)
		if err != nil {
			return nil, err
		}

		closestWarehouse, distance, err := s.DeliveryService.RouteService.CalculateDistance(clientLon, clientLat, warehouses)
		if err != nil {
			return nil, err
		}

		productInfo, err := s.DeliveryService.InventoryService.GetProductInfo(productId)
		if err != nil {
			return nil, err
		}

		deliveryOp, err := s.DeliveryService.GetAvailableDeliveryOptions(closestWarehouse, distance, productInfo)
		if err != nil {
			return nil, err
		}

		productDeliveryOptions = append(productDeliveryOptions, &proto.ProductDeliveryOptions{ProductId: productId, DeliveryOptions: deliveryOp})
	}

	return &proto.DeliveryResponse{
		Products: productDeliveryOptions,
	}, nil
}

func (d *DeliveryService) GetAvailableDeliveryOptions(warehouseInfo *WarehouseInfo, distance float64, productInfo *ProductInfo) ([]*proto.DeliveryOptions, error) {
	var deliveryOptions []*proto.DeliveryOptions

	for _, service := range d.TransportServices {
		available, err := service.CheckAvailability(warehouseInfo.WarehouseID, distance, productInfo.Height, productInfo.Length, productInfo.Width, productInfo.Weight)
		if available && err == nil {
			deliveryOptions = append(deliveryOptions, service.GetDeliveryOption())
			log.Println("Ok")
		} else {
			log.Println("Not Available: ", err)
		}
	}

	return deliveryOptions, nil
}

func (d *DeliveryService) DeliveryInitialization() {}
