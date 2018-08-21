workspace(name = "com_github_stackb_googleapis")

#####################################################################
# RULES_GO
#####################################################################

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "d850f8bbd15d94ce11a078b3933e92ebbf09f715",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.12.0/bazel-gazelle-0.12.0.tar.gz"],
    sha256 = "ddedc7aaeb61f2654d7d7d4fd7940052ea992ccdb031b8f9797ed143ac7e8d43",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

git_repository(
    name = "build_stack_go_github_com_googleapis_googleapis_publish_build_event_6215aac7",
    remote = "https://go.stack.build/github.com/googleapis/googleapis/publish_build_event/6215aac7.git",
    commit = "fd20e52ed40ceeb3066450da48004f9d4ca006ff",
)

load("@build_stack_go_github_com_googleapis_googleapis_publish_build_event_6215aac7//:godeps.bzl", "go_proto_repositories")

go_proto_repositories()

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    urls = ["https://github.com/grpc/grpc-go/archive/1e2570b1b19ade82d8dbb31bba4e65e9f9ef5b34.tar.gz"],
    strip_prefix = "grpc-go-1e2570b1b19ade82d8dbb31bba4e65e9f9ef5b34/",
    build_file_generation = "on",
    build_file_proto_mode = "disable",
)
