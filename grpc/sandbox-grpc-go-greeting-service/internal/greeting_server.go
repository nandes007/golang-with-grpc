package internal

import (
	"context"
	pb "github.com/nandes007/sandbox-grpc-go-greeting-service/helloworld"
	"log"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
