load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "bes.go",
        "main.go",
    ],
    importpath = "github.com/stackb/googleapis/bes",
    visibility = ["//visibility:private"],
    deps = [
        "@build_stack_go_github_com_google_protobuf_empty_6ff2474d//:go_default_library",
        "@build_stack_go_github_com_googleapis_googleapis_publish_build_event_6215aac7//:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//reflection:go_default_library",
    ],
)

go_binary(
    name = "bes",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
