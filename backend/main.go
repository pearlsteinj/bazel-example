package main

import (
	"log"
	"net"
	"os"

	"github.com/aws/aws-xray-sdk-go/xray"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/pearlsteinj/bazel-example/proto/backend"
)

func main() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			xray.UnaryServerInterceptor(
				xray.WithSegmentNamer(
					xray.NewFixedSegmentNamer("hello-world-backend")))))
	server := NewServer()

	pb.RegisterBackendServiceServer(s, server)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	if err := s.Serve(lis); err != nil {
		healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
		log.Fatalf("failed to serve: %v", err)
	}
}
