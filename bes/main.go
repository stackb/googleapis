package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	be "go.stack.build/github.com/googleapis/googleapis/build_events/e8a83bde"
)

func main() {
	pb := &be.BuildEvent{}
	bytes, err := proto.Marshal(pb)
	log.Printf("Hello world! %v %v", bytes, err)
}
