package main

import (
	"log"
	"fmt"
	"net"
	
	proto "go.stack.build/github.com/google/protobuf/at/b4deda09"
	pbe "go.stack.build/github.com/googleapis/googleapis/publish_build_event/6215aac7"
)

func main() {

	sopts := []grpc.ServerOption{grpc.MaxConcurrentStreams(200)}
	grpcServer := grpc.NewServer(sopts...)

	//besService := &build_events.Service{}
	
	//bes.RegisterPublishBuildEventServer(grpcServer, besService)
	log.Printf("Hello world! %v %v", bytes, err)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
			log.Fatalf("failed to listen: %v", err)
	}       
	grpcServer.Serve(lis)

}
