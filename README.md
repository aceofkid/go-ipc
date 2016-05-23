# go-ipc: A library for inter-process communication written in pure Go.
This package gives you access to os-native ipc mechanisms on linux, OSX, freebsd, and windows.
### Features
* Pure Go implementation, no cgo is required.
* Works on linux, OSX, freebsd, and windows.
* Suport of the following mechanisms:
    - fifo (all supported platforms)
    - memory mapped files (all supported platforms)
    - shared memory (all supported platforms)
    - messsage queues (linux, OSX, freebsd)
    - locking primitives: (all supported platforms)

## Install
1. Install Go 1.6 or higher.
2. Run
```
go get -u bitbucket.org/avd/go-ipc
```

## Documentation
Documentation can be found at [`godoc`](https://godoc.com/bitbucket.org/avd/go-ipc).

## Build status
This library is currently alfa. The 'master' branch is not guaranteed to contain stable code,
it is even not guaranteed, that it builds correctly on all platforms. The library uses
[Semantic Versioning 2.0.0](http://semver.org/), so it is recommended to checkout the latest release.

## Contributing
Any contributions are welcome.
Feel free to:

  - create [`issues`](https://bitbucket.org/avd/go-ipc/issues/new)
  - open [`pull requests`](https://bitbucket.org/avd/go-ipc/pull-requests/new)

Before opening a PR, be sure, that:

  - your PR has an issue associated with it.
  - your commit messages are adequate.
  - you added unit tests for your code.
  - you gofmt'ed the code.
  - you used [`gometalinter`](https://github.com/alecthomas/gometalinter) to check your code.

PR's containing documentation improvements and tests are especially welcome.

## LICENSE

This package is made available under an Apache License 2.0. See
LICENSE and NOTICE.
