# Bazel

When we tell Bazel to build something, or when we declare a dependency
or something, we do it by using a "label" to refer to a target

     label = [ [ [ "@" repository ] "//" package ] ":" ] target

 - A "target" is a thing within a package that Bazel can build
 - A "package" is a directory containing a `BUILD` or `BUILD.bazel`
   file.
 - A "repository" is a directory containing a `WORKSPACE` file, the
   `WORKSPACE` file indicates how to find the repositories of its
   dependencies.

# Common commands

Build the Docker images (does NOT load them in to dockerd):
 - `bazel build :ambassador`
 - `bazel build :kat-client`
 - `bazel build :kat-server`

Build the Docker images, and load them in to dockerd:
 - `bazel run :ambassador` (dockerd knows the image as `bazel:ambassador`)
 - `bazel run :kat-client` (dockerd knows the image as `bazel:kat-client`)
 - `bazel run :kat-server` (dockerd knows the image as `bazel:kat-server`)

Build and push the Docker images
 - `bazel run :ambassador.push`
 - `bazel run :kat-client.push`
 - `bazel run :kat-server.push`

Run a command line program on the host:
 - `bazel run //cmd/kubeapply:kubeapply`
 - `bazel run //python:diagd`

Download all dependencies, in order to work offline:
 - `bazel sync`
