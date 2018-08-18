load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/stackb/googleapis/bes",
    visibility = ["//visibility:private"],
    deps = [
        "@build_stack_go_github_com_googleapis_googleapis_build_events_e8a83bde//:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)

go_binary(
    name = "bes",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
