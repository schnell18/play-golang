package main

import (
	"log"
	"net"

	pb "github.com/schnell18/play-golang/grpc/api"
	"github.com/schnell18/play-golang/grpc/internal/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Printf("Server listening at %v\n", lis.Addr())
	pb.RegisterGreeterServer(s, &server.GreeterServerImpl{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
