workspace(name = "com_github_stackb_googleapis")

#####################################################################
# RULES_CLOSURE
#####################################################################

#RULES_CLOSURE_VERSION = "e743ad91efaea4cb57ce474e2b51edcea49ad10a"
RULES_CLOSURE_VERSION = "3555e5ba61fdcc17157dd833eaf7d19b313b1bca"

http_archive(
    name = "io_bazel_rules_closure",
    url = "https://github.com/bazelbuild/rules_closure/archive/%s.zip" % RULES_CLOSURE_VERSION,
    strip_prefix = "rules_closure-%s" % RULES_CLOSURE_VERSION,
)

load("@io_bazel_rules_closure//closure:defs.bzl", "closure_repositories")

closure_repositories()

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

#####################################################################
# PROTO DEPS
#####################################################################

# https://go.stack.build/github.com/googleapis/googleapis/publish_build_event/6215aac7
PUBLISH_BUILD_EVENT_COMMIT = "caafbcb932bd9cf9ab6e4147a79647ee1c618bb7"

# https://go.stack.build/github.com/bazelbuild/bazel/build_event_stream/d6b40d94
BUILD_EVENT_STREAM_COMMIT = "c05218fdea3a272d3d9e1a263196e11f145ee685"

git_repository(
    name = "build_stack_go_github_com_googleapis_googleapis_publish_build_event_6215aac7",
    remote = "https://go.stack.build/github.com/googleapis/googleapis/publish_build_event/6215aac7.git",
    commit = PUBLISH_BUILD_EVENT_COMMIT,
)

load(
    "@build_stack_go_github_com_googleapis_googleapis_publish_build_event_6215aac7//:deps.bzl",
    "publish_build_event_transitive_deps",
)

publish_build_event_transitive_deps()

git_repository(
    name = "build_stack_go_github_com_bazelbuild_bazel_build_event_stream_d6b40d94",
    remote = "https://go.stack.build/github.com/bazelbuild/bazel/build_event_stream/d6b40d94.git",
    commit = BUILD_EVENT_STREAM_COMMIT,
)

load(
    "@build_stack_go_github_com_bazelbuild_bazel_build_event_stream_d6b40d94//:deps.bzl",
    "build_event_stream_transitive_deps",
)

build_event_stream_transitive_deps()

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

http_archive(
    name = "com_github_grpc_grpc",
    urls = ["https://github.com/grpc/grpc/archive/v1.14.1.tar.gz"],
    sha256 = "16f22430210abf92e06626a5a116e114591075e5854ac78f1be8564171658b70",
    strip_prefix = "grpc-1.14.1",
)

load(
    "@com_github_grpc_grpc//bazel:grpc_deps.bzl",
    "grpc_deps",
)
grpc_deps()

git_repository(
    name = "com_github_stackb_grpc_js",
    remote = "https://github.com/stackb/grpc.js.git",
    commit = "bd1a203489b91db158f29c15daf1d213ad179e93",
)
