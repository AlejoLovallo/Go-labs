# Go-backend-examples

Grpc server in go for test purposes.

- Election based on this article: https://medium.com/geekculture/how-to-structure-your-project-in-golang-the-backend-developers-guide-31be05c6fdd9

### Example 1: Simple Ping Pong Server

* [servers/README.md](servers/README.md)

* Docker: 

```bash
docker build -f Dockerfile.server -t go-grpc-server:latest .
docker run -p 50051:50051 go-grpc-server:latest
```

### Example 2: Chat Server

* [chat/README.md](chat/README.md)

