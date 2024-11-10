# Server with channels

### Generate the proto files

```bash
protoc --go_out=. --go-grpc_out=. chat.proto
```

### Running the Server

```bash
go run main.go
```

### Running the Client

* From chat_client folder

```bash
go run main.go
```
