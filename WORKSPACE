load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    strip_prefix = "rules_go-4488af3fdf1cc39c08dde8f76d2da65c8c6553c5",
    urls = ["https://github.com/bazelbuild/rules_go/archive/4488af3fdf1cc39c08dde8f76d2da65c8c6553c5.zip"],
)

http_archive(
    name = "bazel_gazelle",
    strip_prefix = "bazel-gazelle-253128b77088080a348f54d79a28dcd47d99caf9",
    url = "https://github.com/bazelbuild/bazel-gazelle/archive/253128b77088080a348f54d79a28dcd47d99caf9.zip",
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_disintegration_imaging",
    commit = "32df9565b4e0c1460f1915d53f6ff198d9a41af2",
    importpath = "github.com/disintegration/imaging",
)

go_repository(
    name = "org_golang_x_image",
    commit = "991ec62608f3c0da01d400756917825d1e2fd528",
    importpath = "golang.org/x/image",
)
