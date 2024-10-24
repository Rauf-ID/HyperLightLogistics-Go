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
}

type InventoryService struct {
	DB *db.PostgresDB
}

func NewInventoryService(db *db.PostgresDB) *InventoryService {
	return &InventoryService{DB: db}
}

func (s *InventoryService) GetWarehousesForProduct(productId int64) ([]InventoryItem, error) {
	query := `
			SELECT i.product_id, i.warehouse_id, i.quantity, w.location 
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
		err := rows.Scan(&item.ProductID, &item.WarehouseID, &item.Quantity, &item.Location)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		inventoryItems = append(inventoryItems, item)
	}

	return inventoryItems, nil
}
