load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/ypomortsev/imaging-resize
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/ypomortsev/imaging-resize",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_disintegration_imaging//:go_default_library",
        "@org_golang_x_image//font:go_default_library",
        "@org_golang_x_image//font/basicfont:go_default_library",
        "@org_golang_x_image//math/fixed:go_default_library",
    ],
)

go_binary(
    name = "imaging-resize",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
