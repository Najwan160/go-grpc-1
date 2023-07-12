package main

import (
	"context"

	pb "github.com/Najwan160/go-grpc-1/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{
		Message: "Hello",
	}, nil

}
