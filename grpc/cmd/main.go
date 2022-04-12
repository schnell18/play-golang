package main

import (
	"context"
	"log"
	"net"

	pb "github.com/schnell18/play-golang/grpc/api"
	"github.com/schnell18/play-golang/grpc/internal/server"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	log.Printf("Server listening at %v\n", lis.Addr())
	reflection.Register(grpcServer)
	pb.RegisterGreeterServer(grpcServer, &server.GreetServerImpl{})

	// register health check service
	health.RegisterHealthServer(grpcServer, &GrpcHealthCheckService{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

type GrpcHealthCheckService struct {
}

func (s *GrpcHealthCheckService) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}

func (s *GrpcHealthCheckService) Watch(in *health.HealthCheckRequest, _ health.Health_WatchServer) error {
	// watch is left unimplemented
	return nil
}
