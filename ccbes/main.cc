
#include <iostream>
#include <memory>
#include <string>

#include <grpc++/grpc++.h>

#include "google/devtools/build/v1/publish_build_event.pb.h"
#include "google/devtools/build/v1/publish_build_event.grpc.pb.h"

using google::devtools::build::v1::PublishBuildEvent;
using google::devtools::build::v1::PublishBuildEvent;
using google::devtools::build::v1::PublishBuildToolEventStreamRequest;
using google::devtools::build::v1::PublishBuildToolEventStreamResponse;
using google::devtools::build::v1::PublishLifecycleEventRequest;
using google::devtools::build::v1::OrderedBuildEvent;
using google::devtools::build::v1::StreamId;
using google::protobuf::Empty;
using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::ServerReaderWriter;
using grpc::Status;

// Logic and data behind the server's behavior.
class PbeServiceImpl final : public PublishBuildEvent::Service
{
    Status PublishLifecycleEvent(
        ServerContext *context,
        const PublishLifecycleEventRequest *req,
        Empty *resp) override
    {
        std::cout << "PublishLifecycleEvent" << std::endl;
        return Status::OK;
    }

    Status PublishBuildToolEventStream(
        ServerContext *context,
        ServerReaderWriter<PublishBuildToolEventStreamResponse, PublishBuildToolEventStreamRequest>* stream) override
    {
        std::cout << "PublishBuildToolEventStream" << std::endl;
        PublishBuildToolEventStreamRequest req;
        while (stream->Read(&req)) {
            std::cout << "Read" << std::endl;
            OrderedBuildEvent obe = req.ordered_build_event();
            PublishBuildToolEventStreamResponse resp{};
            std::cout << "OrderedBuildEvent" << obe.DebugString() << std::endl;
            StreamId streamId = obe.stream_id();
            //std::cout << "StreamId " << streamId.DebugString() << std::endl;
            //StreamId outId{};
            //outId.set_build_id(streamId.build_id());
            resp.mutable_stream_id()->CopyFrom(streamId);
            std::cout << "Response " << resp.DebugString() << std::endl;
            //resp.set_sequence(obe->get_sequence())
            //std::cout << "Ack " << streamId.sequence_number() << std::endl;
            stream->Write(resp);
        }
        return Status::OK;
    }
};

void RunServer()
{
    std::string server_address("0.0.0.0:50059");
    PbeServiceImpl service;

    ServerBuilder builder;
    // Listen on the given address without any authentication mechanism.
    builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
    // Register "service" as the instance through which we'll communicate with
    // clients. In this case it corresponds to an *synchronous* service.
    builder.RegisterService(&service);
    // Finally assemble the server.
    std::unique_ptr<Server> server(builder.BuildAndStart());
    std::cout << "Server listening on " << server_address << std::endl;

    // Wait for the server to shutdown. Note that some other thread must be
    // responsible for shutting down the server for this call to ever return.
    server->Wait();
}

int main(int argc, char **argv)
{
    RunServer();
    return 0;
}