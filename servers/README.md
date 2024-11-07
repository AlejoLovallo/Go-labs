# GRPC Server

### Folder Structure

- `server.go`: Defines the server implementation.
- `client.go`: Defines the client implementation.
- `main.go`: Entry point for the server.


### Instructions

1. Run `go mod tidy` to install the dependencies.
2. Run `go run main.go` to start the server.
3. Run `go run client.go` to test the server.

protoc --go_out=. --go-grpc_out=. server.proto