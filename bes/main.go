package main

import (
	"fmt"
	"log"
	"net"

	//proto "github.com/golang/protobuf/proto"

	publish_build_event "go.stack.build/github.com/googleapis/googleapis/publish_build_event/6215aac7"

	"google.golang.org/grpc"
)

func main() {
	grpcPort := 50053

	sopts := []grpc.ServerOption{grpc.MaxConcurrentStreams(200)}
	grpcServer := grpc.NewServer(sopts...)

	besService := &BuildEventService{}

	publish_build_event.RegisterPublishBuildEventServer(grpcServer, besService)
	//reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening at :%d", grpcPort)

	endpoints := ListGRPCResources(grpcServer)
	for _, ep := range endpoints {
		log.Println(ep)
	}

	grpcServer.Serve(lis)

}

func ListGRPCResources(server *grpc.Server) []string {
	ret := []string{}
	for serviceName, serviceInfo := range server.GetServiceInfo() {
		for _, methodInfo := range serviceInfo.Methods {
			fullResource := fmt.Sprintf("/%s/%s", serviceName, methodInfo.Name)
			ret = append(ret, fullResource)
		}
	}
	return ret
}
