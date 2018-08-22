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
    commit = "4c01b6771c0631c92380f2ca366bd73655966465",
)

load(
    "@build_stack_go_github_com_googleapis_googleapis_publish_build_event_6215aac7//:godeps.bzl",
    "publish_build_event_proto_repositories",
)

publish_build_event_proto_repositories()

git_repository(
    name = "build_stack_go_github_com_bazelbuild_bazel_build_event_stream_d6b40d94",
    remote = "https://go.stack.build/github.com/bazelbuild/bazel/build_event_stream/d6b40d94.git",
    commit = "9386f2933bc6a51a2ad14d3ab8495134bdd29179",
)

load(
    "@build_stack_go_github_com_bazelbuild_bazel_build_event_stream_d6b40d94//:godeps.bzl",
    "build_event_stream_proto_repositories",
)

build_event_stream_proto_repositories()

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    urls = ["https://github.com/golang/protobuf/archive/bbd03ef6da3a115852eaf24c8a1c46aeb39aa175.tar.gz"],
    strip_prefix = "protobuf-bbd03ef6da3a115852eaf24c8a1c46aeb39aa175/",
    build_file_generation = "on",
    build_file_proto_mode = "disable",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    urls = ["https://github.com/grpc/grpc-go/archive/1e2570b1b19ade82d8dbb31bba4e65e9f9ef5b34.tar.gz"],
    strip_prefix = "grpc-go-1e2570b1b19ade82d8dbb31bba4e65e9f9ef5b34/",
    build_file_generation = "on",
    build_file_proto_mode = "disable",
)
