package main

import (
	"context"
	"log"
	"net"

	"github.com/Sohail-9098/vehicle-data-collector/internal/protobufs/vehicle"
	"google.golang.org/grpc"
)

type server struct {
	vehicle.UnimplementedDataProcessingServiceServer
}

func (s *server) ProcessTelemetryData(ctx context.Context, data *vehicle.Telemetry) (*vehicle.Empty, error) {
	log.Println("received data: ", data)
	return &vehicle.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port: 50051 error: %v", err)
	}
	grpcServer := grpc.NewServer()
	vehicle.RegisterDataProcessingServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
