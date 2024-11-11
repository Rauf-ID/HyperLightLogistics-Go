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

import "HyperLightLogistics-Go/internal/service/transport"

type DeliveryService struct {
	DroneService *transport.DroneService
}

type DeliveryOption struct {
	Type         string
	DeliveryTime string
	Price        float64
}

func NewDeliveryService(droneService *transport.DroneService) *DeliveryService {
	return &DeliveryService{
		DroneService: droneService,
	}
}

func (d *DeliveryService) GetAvailableDeliveryOptions(distance float64, productInfo *ProductInfo) ([]DeliveryOption, error) {
	d.DroneService.CheckDroneAvailability(distance, productInfo.Height, productInfo.Length, productInfo.Width, productInfo.Weight)

	return nil, nil
}