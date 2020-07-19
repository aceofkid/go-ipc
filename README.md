# go-ipc: A library for inter-process communication written in pure Go.
This package gives you access to os-native ipc mechanisms on Linux, OSX, FreeBSD, and Windows.

[![CircleCI](https://circleci.com/bb/aceofkid/go-ipc/tree/master.svg?style=svg)](https://circleci.com/bb/aceofkid/go-ipc/tree/master) [![GoDoc](https://godoc.org/github.com/aceofkid/go-ipc?status.svg)](https://godoc.org/github.com/aceofkid/go-ipc) [![Go Report Card](https://goreportcard.com/badge/github.com/aceofkid/go-ipc)](https://goreportcard.com/report/github.com/aceofkid/go-ipc) 


* Pure Go implementation, no cgo is required.
* Works on Linux, OSX, FreeBSD, and Windows (x86 or x86-64).
* Support of the following mechanisms:
    - fifo (unix and windows pipes)
    - memory mapped files
    - shared memory
    - system message queues (Linux, FreeBSD, OSX)
    - cross-platform priority message queue
    - mutexes, rw mutexes
    - semaphores
    - events
    - conditional variables

## Install
1. Install Go 1.4 or higher.
2. Run
```
go get -u github.com/aceofkid/go-ipc
```

## System requirements
1. Linux, OSX, FreeBSD, and Windows (x86 or x86-64).
2. Go 1.4 or higher.

## Documentation
Documentation can be found at [`godoc`](https://godoc.org/github.com/aceofkid/go-ipc).

## Notes

## Build status
This library is currently beta. The 'master' branch is not guaranteed to contain stable code,
it is even not guaranteed, that it builds correctly on all platforms. The library uses
[Semantic Versioning 2.0.0](http://semver.org/), so it is recommended to checkout the latest release.

## Contributing
Any contributions are welcome.
Feel free to:

  - create [`issues`](https://github.com/aceofkid/go-ipc/issues/new)
  - open [`pull requests`](https://github.com/aceofkid/go-ipc/pull-requests/new)

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
