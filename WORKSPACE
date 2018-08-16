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

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

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
