// Client to call the server

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/AlejoLovallo/Go-Grpc/servers/proto/servers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up connection to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create client
	client := pb.NewPingServiceClient(conn)

	// Contact the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call Ping RPC
	resp, err := client.Ping(ctx, &pb.PingRequest{Message: "Hello from client!"})
	if err != nil {
		log.Fatalf("Could not ping: %v", err)
	}

	fmt.Printf("Response from server: %s\n", resp.GetMessage())
	fmt.Printf("Timestamp: %s\n", resp.GetTimestamp())
}
