# go-example-package

This repository is used for experimenting with `go` packaging. Packaging in `go` is different
to Python or C++ which most users are familiar with. We use a subset of the conventional
structure found in Go's recommended
[project layout](https://github.com/golang-standards/project-layout)

## Contents

- `/pkg`
    - The backend of the package, where all the inner workings
lie
- `/internal`
    - Private modules and packages - this isn't particularly
    important and currently has no files.
- `/cmd`
    - A place to store the files used to build executables,
typically the files here import things from `/pkg`
- `/test`
    - Tests go here, which follow standard `go` testing
- `/.github/workflows/build.yml`
    - Build and Test the whole package, look here for commands
    to do this

## Contributing

- Ensure everything you add is related to string manipulation
- This is because string manipulation is easy to understand, hence
the code can be implicitly understood
- Ensure all tests work and any new things are tested - this is just
good practice for any coding that is done ever