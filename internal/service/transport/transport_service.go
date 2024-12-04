package transport

import "HyperLightLogistics-Go/internal/service/proto"

type TransportService interface {
	CheckAvailability(warehouseID int64, distance float64, height, length, width, weight float32) (bool, error)
	GetDeliveryOption() *proto.DeliveryOptions
}
