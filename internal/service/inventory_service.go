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
	"HyperLightLogistics-Go/internal/db"
	"context"
	"log"
)

type InventoryItem struct {
	ProductID   int64
	WarehouseID int64
	Quantity    int64
	Location    string
	Latitude    float32
	Longitude   float32
}

type InventoryService struct {
	DB *db.PostgresDB
}

func NewInventoryService(db *db.PostgresDB) *InventoryService {
	return &InventoryService{DB: db}
}

func (s *InventoryService) GetWarehousesForProduct(productId uint64) ([]InventoryItem, error) {
	query := `
			SELECT i.product_id, i.warehouse_id, i.quantity, w.location, w.latitude, w.longitude
			FROM inventory as i, warehouses as w
			WHERE w.id = i.warehouse_id and i.product_id = $1
	`

	rows, err := s.DB.Conn.Query(context.Background(), query, productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventoryItems []InventoryItem

	for rows.Next() {
		var item InventoryItem
		err := rows.Scan(&item.ProductID, &item.WarehouseID, &item.Quantity, &item.Location, &item.Latitude, &item.Longitude)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		inventoryItems = append(inventoryItems, item)
	}

	return inventoryItems, nil
}
