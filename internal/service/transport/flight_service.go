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

package transport

import (
	"HyperLightLogistics-Go/internal/service/proto"
	"errors"
)

type FlightService struct {
}

func NewFlightService() *FlightService {
	return &FlightService{}
}

func (d *FlightService) CheckAvailability(distance float64, height, length, width, weight float32) (bool, error) {
	if distance >= 50000 || height > 5.0 || length > 5.0 || width > 5.0 || weight > 10.0 {
		return false, errors.New("product exceeds allowable drone limits for distance, size, or weight")
	}
	return true, nil
}

func (d *FlightService) GetDeliveryOption() *proto.DeliveryOptions {
	return &proto.DeliveryOptions{
		Type:         "Flight Delivery",
		Price:        1.5,
		DeliveryTime: "8:00 AM to 3:00 PM",
	}
}
