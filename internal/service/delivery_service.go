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
	"log"
)

type DeliveryService struct {
	DroneService  *transport.DroneService
	VanService    *transport.VanService
	TruckService  *transport.TruckService
	FlightService *transport.FlightService
}

func NewDeliveryService(droneService *transport.DroneService, vanService *transport.VanService,
	truckService *transport.TruckService, flightService *transport.FlightService) *DeliveryService {
	return &DeliveryService{
		DroneService:  droneService,
		VanService:    vanService,
		TruckService:  truckService,
		FlightService: flightService,
	}
}

func (d *DeliveryService) GetAvailableDeliveryOptions(distance float64, productInfo *ProductInfo) ([]*proto.DeliveryOptions, error) {
	var deliveryOptions []*proto.DeliveryOptions

	droneAvailable, err := d.DroneService.CheckDroneAvailability(distance, productInfo.Height, productInfo.Length, productInfo.Width, productInfo.Weight)
	if droneAvailable && err == nil {
		log.Println("Ok")
		deliveryOptions = append(deliveryOptions, &proto.DeliveryOptions{Type: "Drone Delivery", Price: 1.5, DeliveryTime: "8:00 AM to 3:00 PM"})
	} else {
		log.Println("Not Ok: ", err)
	}

	vanAvailable, err := d.VanService.CheckVanAvailability(distance, productInfo.Height, productInfo.Length, productInfo.Width, productInfo.Weight)
	if vanAvailable && err == nil {
		log.Println("Ok")
		deliveryOptions = append(deliveryOptions, &proto.DeliveryOptions{Type: "Van Delivery", Price: 1.5, DeliveryTime: "8:00 AM to 3:00 PM"})
	} else {
		log.Println("Not Ok: ", err)
	}

	if distance >= 16000 {
		truckAvailable, err := d.TruckService.CheckTruckAvailability(distance, productInfo.Height, productInfo.Length, productInfo.Width, productInfo.Weight)
		if truckAvailable && err == nil {
			log.Println("Ok")
			deliveryOptions = append(deliveryOptions, &proto.DeliveryOptions{Type: "Truck Delivery", Price: 1.5, DeliveryTime: "8:00 AM to 3:00 PM"})
		} else {
			log.Println("Not Ok: ", err)
		}

		flightAvailable, err := d.FlightService.CheckFlightAvailability(distance, productInfo.Height, productInfo.Length, productInfo.Width, productInfo.Weight)
		if flightAvailable && err == nil {
			log.Println("Ok")
			deliveryOptions = append(deliveryOptions, &proto.DeliveryOptions{Type: "Flight Delivery", Price: 1.5, DeliveryTime: "8:00 AM to 3:00 PM"})
		} else {
			log.Println("Not Ok: ", err)
		}
	}

	return deliveryOptions, nil
}
