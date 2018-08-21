package main

import (
	"fmt"
	"log"
	"net"

	//proto "github.com/golang/protobuf/proto"

	"google.golang.org/grpc"
)

func main() {
	grpcPort := 50051

	sopts := []grpc.ServerOption{grpc.MaxConcurrentStreams(200)}
	grpcServer := grpc.NewServer(sopts...)

	//besService := &build_events.Service{}

	//bes.RegisterPublishBuildEventServer(grpcServer, besService)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)

}
