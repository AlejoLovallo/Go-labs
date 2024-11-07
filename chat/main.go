package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/AlejoLovallo/Go-Grpc/chat/proto/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// server implements the ChatService server
type server struct {
	pb.UnimplementedChatServiceServer
}

// ChatStream handles the bidirectional streaming RPC
func (s *server) ChatStream(stream pb.ChatService_ChatStreamServer) error {
	for {
		// Receive message from client
		msg, err := stream.Recv()
		if err == io.EOF {
			// Client closed the stream
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving message from client: %v", err)
			return err
		}

		// Get client information
		p, ok := peer.FromContext(stream.Context())
		if ok {
			log.Printf("Message received from %s: %s", msg.GetUser(), msg.GetMessage())
			log.Printf("Client address: %v", p.Addr)
		}

		// Respond back to the client
		response := &pb.ChatMessage{
			User:    "Server",
			Message: fmt.Sprintf("Message received from %s: %s", msg.GetUser(), msg.GetMessage()),
		}

		if err := stream.Send(response); err != nil {
			log.Fatalf("Error sending message to client: %v", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})

	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
