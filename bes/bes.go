package main

import (
	"context"
	"io"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	//any "go.stack.build/github.com/google/protobuf/any/b3a0960f"
	bes "go.stack.build/github.com/bazelbuild/bazel/build_event_stream/d6b40d94"
	empty "go.stack.build/github.com/google/protobuf/empty/6ff2474d"
	be "go.stack.build/github.com/googleapis/googleapis/build_events/e8a83bde"
	pbe "go.stack.build/github.com/googleapis/googleapis/publish_build_event/6215aac7"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type BuildEventService struct {
	//LifecycleEventQueue map[string]*chan *pbe.PublishLifecycleEventRequest
	//BuildEventQueue map[string]chan *pbe.OrderedBuildEvent
	//LifecycleQueue chan *pbe.PublishLifecycleEventRequest
	//ToolQueue      chan *pbe.OrderedBuildEvent
}

func NewBuildEventService() *BuildEventService {
	return &BuildEventService{}
}

func (s *BuildEventService) PublishLifecycleEvent(ctx context.Context, req *pbe.PublishLifecycleEventRequest) (*empty.Empty, error) {
	log.Println("Got publish lifecycle event %+v", req)
	key := req.ProjectId
	//queue := s.LifecycleQueue
	log.Printf("Got publish event %s...", key)
	//queue <- req
	//log.Printf("lce sent to channel %s\n", key)
	return &empty.Empty{}, nil
}

// Publish build tool events belonging to the same stream to a backend job
// using bidirectional streaming.
func (s *BuildEventService) PublishBuildToolEventStream(stream pbe.PublishBuildEvent_PublishBuildToolEventStreamServer) error {
	log.Println("Got tool event stream message(s)")
	graph := NewBuildEventGraph()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("DONE")
			return nil
		}
		if err != nil {
			return err
		}
		obe := req.OrderedBuildEvent
		streamId := obe.StreamId
		seqNum := obe.SequenceNumber
		event := obe.Event

		if event != nil {
			log.Printf("Event: %s", event.EventTime)
			err := handleEvent(graph, event)
			if err != nil {
				log.Printf("Event error: %v", err)
				return grpc.Errorf(codes.InvalidArgument, "%v", err)
			}
		}

		//log.Println("Writing tool event to channel... +%v", event);
		log.Printf("Commit %s %s", streamId, seqNum)
		stream.Send(&pbe.PublishBuildToolEventStreamResponse{
			StreamId:       streamId,
			SequenceNumber: seqNum,
		})
	}
	return nil
}

func handleEvent(listener BuildEventListener, event *be.BuildEvent) error {
	if e := event.GetBuildEnqueued(); e != nil {
		log.Printf("Build Enqueued")
	} else if e := event.GetBuildFinished(); e != nil {
		log.Printf("Build Finished")
	} else if e := event.GetInvocationAttemptStarted(); e != nil {
		log.Printf("Invocation Attempt Started")
	} else if e := event.GetInvocationAttemptFinished(); e != nil {
		log.Printf("Invocation Attempt Finished")
	} else if e := event.GetConsoleOutput(); e != nil {
		log.Printf("Console Output")
	} else if e := event.GetComponentStreamFinished(); e != nil {
		log.Printf("Component Stream Finished")
	} else if e := event.GetBuildExecutionEvent(); e != nil {
		log.Printf("Build Execution Event")
	} else if e := event.GetBazelEvent(); e != nil {
		log.Printf("Bazel Event")
		buildEvent, err := unmarshalBazelEvent(event)
		if err != nil {
			log.Printf("Unhandled bazel event: %v (%v)", e, err)
		} else {
			return handleBazelEvent(listener, buildEvent)
		}
	} else if e := event.GetSourceFetchEvent(); e != nil {
		log.Printf("Source Fetch Event")
	} else {
		log.Printf("Unhandled type: %s", event)
	}
	return nil
}

func handleBazelEvent(listener BuildEventListener, event *bes.BuildEvent) error {
	NotifyBuildEvent(event, listener)
	return nil
}

// unmarshalBazelEvent remarshals the any proto across the type
// boundary that exists since we have 2 any proto definitions.
func unmarshalBazelEvent(e *be.BuildEvent) (*bes.BuildEvent, error) {
	var buildEvent bes.BuildEvent
	err := proto.Unmarshal(e.GetBazelEvent().GetValue(), &buildEvent)
	if err != nil {
		return nil, err
	}
	return &buildEvent, nil
}

// unmarshalBazelEvent remarshals the any proto across the type
// boundary that exists since we have 2 any proto definitions.
func unmarshalBazelEventOld(e *be.BuildEvent) (*bes.BuildEvent, error) {
	var buildEvent bes.BuildEvent
	bytes, err := proto.Marshal(e)
	if err != nil {
		return nil, err
	}
	log.Printf("Marshalled any; %v", bytes)
	var anyProto any.Any
	err = proto.Unmarshal(bytes, &anyProto)
	if err != nil {
		return nil, err
	}
	log.Printf("Unmarshalled any; %v", anyProto)
	err = ptypes.UnmarshalAny(&anyProto, &buildEvent)
	if err != nil {
		return nil, err
	}
	return &buildEvent, nil
}
