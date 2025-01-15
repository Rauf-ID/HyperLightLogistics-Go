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
	"HyperLightLogistics-Go/internal/db"
	"HyperLightLogistics-Go/internal/service/proto"
	"context"
	"errors"
	"log"
	"time"
)

type TruckService struct {
	DB *db.PostgresDB
}

func NewTruckService(db *db.PostgresDB) *TruckService {
	return &TruckService{DB: db}
}

func (d *TruckService) CheckAvailability(warehouseID int64, distance float64, height, length, width, weight float32) (bool, error) {
	if distance <= 50000 || height > 5.0 || length > 5.0 || width > 5.0 || weight > 10.0 {
		return false, errors.New("product exceeds allowable truck limits for distance, size, or weight")
	}

	query := `
		SELECT id, capacity, load, next_available_time, status
		FROM trucks
		WHERE warehouse_id = $1 and status = 'available' and capacity >= $2 and next_available_time <= $3
		ORDER BY next_available_time ASC
		LIMIT 1
	`

	now := time.Now()
	row := d.DB.Conn.QueryRow(context.Background(), query, warehouseID, weight, now)

	var id int64
	var capacity, load float32
	var nextAvailableTime time.Time
	var status string

	err := row.Scan(&id, &capacity, &load, &nextAvailableTime, &status)
	if err != nil {
		log.Println("No suitable trucks available:", err)
		return false, errors.New("no suitable trucks available")
	}

	log.Printf("Truck %d is available with capacity %.2f\n", id, capacity)
	return true, nil
}

func (d *TruckService) GetDeliveryOption() *proto.DeliveryOptions {
	return &proto.DeliveryOptions{
		Type:         "Truck Delivery",
		Price:        1.5,
		DeliveryTime: "8:00 AM to 3:00 PM",
	}
}
