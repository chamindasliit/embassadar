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

 - `bazel build :ambassador`
 - `bazel build :kat-client`
 - `bazel build :kat-server`

 - `bazel run :ambassador.push`
 - `bazel run :kat-client.push`
 - `bazel run :kat-server.push`

 - `bazel run //cmd/kubeapply:kubeapply`
 - `bazel run @ambassador_python//ambassador_diag:diagd`
