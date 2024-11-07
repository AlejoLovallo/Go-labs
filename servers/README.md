# GRPC Server

This is a simple gRPC server implementation in Go that provides a ping service. The server listens on port 50051 and responds to ping requests with a "Pong" message and timestamp.

### Features
- Basic ping/pong service
- Timestamp included in response
- TCP server on port 50051

### Generate the proto files

```
protoc --go_out=. --go-grpc_out=. server.proto
```

### Running the Server

1. Start the server:

* From root folder

```bash
go run main.go
```

2. Start the client

* From server_client folder

```bash
go run main.go
```


