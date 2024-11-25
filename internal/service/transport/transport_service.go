package transport

import "HyperLightLogistics-Go/internal/service/proto"

type TransportService interface {
	CheckAvailability(distance float64, height, length, width, weight float32) (bool, error)
	GetDeliveryOption() *proto.DeliveryOptions
}
