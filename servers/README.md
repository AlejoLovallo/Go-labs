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


### Docker

* Create network

```bash
docker network create grpc-network
```

* Build docker server image

```bash
docker build -f Dockerfile.server_client -t go-grpc-client:latest . 
```

* Build docker client image

```bash
docker build -f Dockerfile.server_client -t go-grpc-client:latest . 
```

* Run docker server

```bash
docker run -d --name grpc-server --network grpc-server-network -p 50051:50051 go-grpc-server:latest
```

* Run docker client

```bash
docker run --name grpc-server-client --network grpc-server-network go-grpc-client:latest
```

* Notice the change on the ip address on the client side:

```go
conn, err := grpc.Dial("grpc-server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
```


### Using docker compose

* Build 

```bash
docker compose -f docker-compose.server.yaml build
```

* Run

```bash
 docker compose -f docker-compose.server.yaml up -d  
```

* Here the change will as the name of the service server

```go
conn, err := grpc.Dial("server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
```
