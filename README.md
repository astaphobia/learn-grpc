# Learn gRPC
A learning experiment with gRPC based upon the [official documentation](http://www.grpc.io/docs/) and a [Medium article](https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b) by [Shiju Varghese](https://medium.com/@shijuvar)

## About
This repo contains a protocol buffer, for a customer service, as defined in the [Medium article](https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b). This single protocol buffer has then been implemented in 4 different languages:

* [Golang](https://github.com/danmrichards/learn-grpc/tree/master/golang)
* [NodeJS](https://github.com/danmrichards/learn-grpc/tree/master/nodejs)
* [PHP](https://github.com/danmrichards/learn-grpc/tree/master/php)
* [Python](https://github.com/danmrichards/learn-grpc/tree/master/python)

Each of these languages has a client and server implementation (with the exception of PHP which just has a client). The implementations all use generated client/server libraries from the (protoc)[https://github.com/google/protobuf] tool.

## Usage
See the README in each language folder for more information.

> Hint: Yeah you could be dull and run the client and server in the same language. But you don't have to! The magic of protocol buffers means the client/server implementation doesn't matter. It just works.
