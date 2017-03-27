# gRPC - Python

## Prerequisites

Ensure you have `pip` version 8 or higher:

```bash
$ python -m pip install --upgrade pip
```

If you cannot upgrade `pip` due to a system-owned installation, you can
run the example in a virtualenv:

```bash
$ python -m pip install virtualenv
$ virtualenv venv
$ source venv/bin/activate
$ python -m pip install --upgrade pip
```

### Install gRPC

Install gRPC:

```bash
$ python -m pip install grpcio
```

Or, to install it system wide:

```bash
$ sudo python -m pip install grpcio
```

On El Capitan OSX, you may get the following error:

```bash
$ OSError: [Errno 1] Operation not permitted: '/tmp/pip-qwTLbI-uninstall/System/Library/Frameworks/Python.framework/Versions/2.7/Extras/lib/python/six-1.4.1-py2.7.egg-info'
```

You can work around this using:

```bash
$ python -m pip install grpcio --ignore-installed
```

### Install gRPC tools

Python's gRPC tools include the protocol buffer compiler `protoc` and the
special plugin for generating server and client code from `.proto` service
definitions. For the first part of our quickstart example, we've already
generated the server and client stubs from
[helloworld.proto](https://github.com/grpc/grpc/tree/{{site.data.config.grpc_release_branch}}/examples/protos/helloworld.proto),
but you'll need the tools for the rest of our quickstart, as well as later
tutorials and your own projects.

To install gRPC tools, run:

```bash
$ python -m pip install grpcio-tools
```

## Server
```bash
$ python server.py
```

## Client
```bash
$ python client.py
```
