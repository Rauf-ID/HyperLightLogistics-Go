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

type FlightService struct {
	DB *db.PostgresDB
}

func NewFlightService(db *db.PostgresDB) *FlightService {
	return &FlightService{DB: db}
}

func (d *FlightService) CheckAvailability(warehouseID int64, distance float64, height, length, width, weight float32) (bool, error) {
	if distance <= 50000 || height > 5.0 || length > 5.0 || width > 5.0 || weight > 10.0 {
		return false, errors.New("product exceeds allowable flight limits for distance, size, or weight")
	}

	query := `
		SELECT id, origin_airport_id, destination_airport_id, status, departure_time, arrival_time, capacity, load
		FROM flights
		WHERE status = 'available' and capacity >= $1 and departure_time > $2
		ORDER BY next_available_time ASC
		LIMIT 1
	`

	now := time.Now()
	row := d.DB.Conn.QueryRow(context.Background(), query, weight, now)

	var id int64
	var capacity, load float32
	var arrival_time, departure_time time.Time
	var origin_airport_id, destination_airport_id, status string

	err := row.Scan(&id, &origin_airport_id, &destination_airport_id, &status, &departure_time, &arrival_time, &capacity, &load)
	if err != nil {
		log.Println("No suitable flights available:", err)
		return false, errors.New("no suitable flights available")
	}

	log.Printf("Flight %d is available with capacity %.2f\n", id, capacity)
	return true, nil
}

func (d *FlightService) GetDeliveryOption() *proto.DeliveryOptions {
	return &proto.DeliveryOptions{
		Type:         "Flight Delivery",
		Price:        1.5,
		DeliveryTime: "8:00 AM to 3:00 PM",
	}
}
