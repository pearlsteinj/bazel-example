package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/pearlsteinj/bazel-example/proto/backend"
)

type Server struct {
	pb.UnimplementedBackendServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) HelloWorld(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	if in.GetRequestId() > 69 {
		return nil, status.Error(codes.InvalidArgument, "Request ID too high")
	}

	return &pb.HelloWorldResponse{
		ResponseId:      uuid.New().String(),
		MessageResponse: "Hello, World!",
	}, nil
}
