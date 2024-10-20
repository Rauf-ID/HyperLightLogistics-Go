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
}

type InventoryService struct {
	DB *db.PostgresDB
}

func NewInventoryService(db *db.PostgresDB) *InventoryService {
	return &InventoryService{DB: db}
}

func (s *InventoryService) GetWarehousesForProduct(productId int64) ([]InventoryItem, error) {
	query := `SELECT product_id, warehouse_id, quantity FROM inventory WHERE product_id = $1`

	rows, err := s.DB.Conn.Query(context.Background(), query, productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventoryItems []InventoryItem

	for rows.Next() {
		var item InventoryItem
		err := rows.Scan(&item.ProductID, &item.WarehouseID, &item.Quantity)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		inventoryItems = append(inventoryItems, item)
	}

	return inventoryItems, nil
}
