package servers

import (
	"context"
	"log"
	pb "proto/servers"
	"time"
)

type server struct {
	pb.UnimplementedPingServiceServer
}

func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received ping request: %v", req.Message)

	return &pb.PingResponse{
		Message:   "Pong: " + req.Message,
		Timestamp: time.Now().Format(time.RFC3339),
	}, nil

}
