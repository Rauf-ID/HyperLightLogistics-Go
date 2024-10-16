package main

import (
	proto "HyperLightLogistics-Go/internal/service"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type DeliveryOptionsServer struct {
	proto.UnimplementedDeliveryOptionsServiceServer
}

func (s DeliveryOptionsServer) CalculateDeliveryOptions(context.Context, *proto.DeliveryRequest) (*proto.DeliveryResponse, error) {

	deliveryOptions := calculateDeliveryRoutes()

	return &proto.DeliveryResponse{
		DeliveryOptions: deliveryOptions,
	}, nil
}

func calculateDeliveryRoutes() []*proto.DeliveryOptions {
	var deliveryOptions []*proto.DeliveryOptions

	deliveryOptions = append(deliveryOptions, &proto.DeliveryOptions{
		Type:         "Standard",
		DeliveryTime: "3-5 days",
		Price:        5.99,
	})

	deliveryOptions = append(deliveryOptions, &proto.DeliveryOptions{
		Type:         "Expedited",
		DeliveryTime: "1-3 days",
		Price:        15.99,
	})

	return deliveryOptions
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &DeliveryOptionsServer{}
	proto.RegisterDeliveryOptionsServiceServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
