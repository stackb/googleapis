load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "bes.go",
        "graph.go",
        "main.go",
    ],
    importpath = "github.com/stackb/googleapis/bes",
    visibility = ["//visibility:private"],
    deps = [
        "@build_stack_go_github_com_bazelbuild_bazel_build_event_stream_d6b40d94//go",  # keep
        "@build_stack_go_github_com_google_protobuf_empty_6ff2474d//go",  # keep
        "@build_stack_go_github_com_googleapis_googleapis_build_events_e8a83bde//go",  # keep
        "@build_stack_go_github_com_googleapis_googleapis_publish_build_event_6215aac7//go",  # keep
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        "@io_bazel_rules_go//proto/wkt:any_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes:go_default_library",
    ],
)

go_binary(
    name = "bes",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
