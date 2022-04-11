package server

import pb "github.com/schnell18/play-golang/grpc/api"
import "context"

type GreetServerImpl struct {
	pb.UnimplementedGreeterServer
}

func (s *GreetServerImpl) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello " + req.GetName(),
	}, nil
}
