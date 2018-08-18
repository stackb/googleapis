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

# # http_archive(
# #     name = "com_github_scele_rules_go_dep",
# #     urls = ["https://github.com/scele/rules_go_dep/archive/49a5e4ca9f6a16c9b4c930a51ce3a537498bb4e1.tar.gz"],
# #     strip_prefix = "rules_go_dep-49a5e4ca9f6a16c9b4c930a51ce3a537498bb4e1",
# #     sha256 = "f170d3d6f55e216f1493f975cde6c489d7070da2a8a41fd4de9812d96f4fb38b",
# # )

# local_repository(
#     name = "com_github_scele_rules_go_dep",
#     path = "/home/pcj/go/src/github.com/pcj/rules_go_dep",
# )

# load("@com_github_scele_rules_go_dep//dep:dep.bzl", "dep_import")

# # https://gist.github.com/subfuzion/12342599e26f5094e4e2d08e9d4ad50d
# dep_import(
#     name = "godeps",
#     prefix = "github.com/stackb/googleapis",
#     gopkg_lock = "//:Gopkg.lock",
# )

# load("@godeps//:Gopkg.bzl", "go_deps")

# go_deps()

# gazelle:repo bazel_gazelle

go_repository(
    name = "build_stack_go_github_com_google_protobuf_any_b3a0960f",
    importpath = "go.stack.build/github.com/google/protobuf/any/b3a0960f",
    commit = "12e56ccdf5b60cb22dd9ee6fbf1540a6d284b691",
)

go_repository(
    name = "build_stack_go_github_com_google_protobuf_descriptor_713ee215",
    importpath = "go.stack.build/github.com/google/protobuf/descriptor/713ee215",
    commit = "7dd38495dc455a7de56056295898fde6761edc37",
)

go_repository(
    name = "build_stack_go_github_com_google_protobuf_timestamp_dafad809",
    importpath = "go.stack.build/github.com/google/protobuf/timestamp/dafad809",
    commit = "477878dc7acaef8814510213a46123631cc60d60",
)

go_repository(
    name = "build_stack_go_github_com_google_protobuf_wrappers_ce39ab37",
    importpath = "go.stack.build/github.com/google/protobuf/wrappers/ce39ab37",
    commit = "7b4c9c626c3225ccf9631c775be07f8c64801307",
)

go_repository(
    name = "build_stack_go_github_com_googleapis_googleapis_status_cfe88cab",
    importpath = "go.stack.build/github.com/googleapis/googleapis/status/cfe88cab",
    commit = "d67f6bc2dbf4c505c821afeac4fe4c89fff7ea94",
)

go_repository(
    name = "build_stack_go_github_com_googleapis_googleapis_http_2c560493",
    importpath = "go.stack.build/github.com/googleapis/googleapis/http/2c560493",
    commit = "45937d74db365dbb049ae97e0930ac0642ed3444",
)

go_repository(
    name = "build_stack_go_github_com_googleapis_googleapis_annotations_164b695e",
    importpath = "go.stack.build/github.com/googleapis/googleapis/annotations/164b695e",
    commit = "d8f543a74582189dc7f216a184e64be4c814d472",
)

go_repository(
    name = "build_stack_go_github_com_googleapis_googleapis_build_status_2b963b92",
    importpath = "go.stack.build/github.com/googleapis/googleapis/build_status/2b963b92",
    commit = "48c4dd8296e5a7751f4a05b1b067c730f837fba2",
)

go_repository(
    name = "build_stack_go_github_com_googleapis_googleapis_build_events_e8a83bde",
    importpath = "go.stack.build/github.com/googleapis/googleapis/build_events/e8a83bde",
    commit = "7cd50fa0bec58a84ba5cb8016a9629d5263f68b6",
)
