package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	pb "github.com/AlejoLovallo/Go-Grpc/chat/proto/chat"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	// Create a bidirectional stream
	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("Error opening stream: %v", err)
	}

	// Channel for sending messages
	done := make(chan struct{})

	// Go routine for receiving messages from the server
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive a message: %v", err)
			}
			log.Printf("Server: %s", in.GetMessage())
		}
		close(done)
	}()

	// Go routine for sending messages to the server
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			log.Print("Enter message: ")
			if scanner.Scan() {
				msg := scanner.Text()
				if err := stream.Send(&pb.ChatMessage{User: "Client", Message: msg}); err != nil {
					log.Fatalf("Failed to send a message: %v", err)
				}
			}
		}
	}()

	// Wait for the server to close the connection
	<-done
	stream.CloseSend()
}
