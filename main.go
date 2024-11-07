// File to run the server
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/AlejoLovallo/Go-Grpc/servers/proto/servers"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPingServiceServer
}

func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received ping request: %v", req.Message)
	return &pb.PingResponse{
		Message:   "Pong: " + req.Message,
		Timestamp: time.Now().String(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPingServiceServer(s, &server{})

	fmt.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
