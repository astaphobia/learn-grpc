# gRPC - Golang

## Prerequisites

### Go version

gRPC works with Go 1.5 or higher.

```bash
$ go version
```

For installation instructions, follow this guide: [Getting Started - The Go Programming Language](https://golang.org/doc/install)

### Install gRPC

Use the following command to install gRPC.

```bash
$ go get google.golang.org/grpc
```

### Install Protocol Buffers v3

Install the protoc compiler that is used to generate gRPC service code. The simplest way to do this is to download pre-compiled binaries for your platform(`protoc-<version>-<platform>.zip`) from here: [https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)

  * Unzip this file.
  * Update the environment variable `PATH` to include the path to the protoc binary file.

Next, install the protoc plugin for Go

```
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

The compiler plugin, protoc-gen-go, will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler, protoc, to find it.  

```
$ export PATH=$PATH:$GOPATH/bin
```

## Server
```bash
$ go run client/main.go
```

## Client
```bash
$ go run server/main.go
```
