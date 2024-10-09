package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"time"

	"github.com/TrapLord92/gRPC-Microservices-in-Go.git/Grpc4BidiStreaming/bidipb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a connection to the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the BidiService
	client := bidipb.NewBidiServiceClient(conn)

	// Create a bidirectional stream
	stream, err := client.BidiChat(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	// Run a goroutine to handle receiving messages from the server
	go func() {
		for {
			// Receive messages from the server
			res, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error while receiving response: %v", err)
				return
			}
			log.Printf("Received from server: %s", res.GetReply())
		}
	}()

	// Simulate sending messages from the client to the server
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Read input from the user
		log.Print("Enter message: ")
		scanner.Scan()
		text := scanner.Text()

		// Send the message to the server
		err := stream.Send(&bidipb.BidiRequest{
			Message: text,
		})
		if err != nil {
			log.Fatalf("Error while sending message: %v", err)
			return
		}

		// Simulate some delay between messages
		time.Sleep(1 * time.Second)
	}
}
