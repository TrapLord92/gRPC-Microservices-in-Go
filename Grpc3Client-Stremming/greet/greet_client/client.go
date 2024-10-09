package main

import (
	"context"
	"log"
	"time"

	"github.com/TrapLord92/gRPC-Microservices-in-Go.git/Grpc3Client-Stremming/greet/greetpb/clientstremmingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a connection to the gRPC server
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials())) // Use insecure for non-TLS
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a client for the GreetClientStreamService
	client := clientstremmingpb.NewGreetClientStreamServiceClient(conn)

	// Call the GreetStreeming method and get the stream
	stream, err := client.GreetStreeming(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetStreeming: %v", err)
	}

	// Send multiple requests as part of the client-side streaming
	requests := []*clientstremmingpb.GreetClientStreamRequest{
		{GreetingClient: &clientstremmingpb.GreetingClientStream{FirstName: "John", LastName: "Doe"}},
		{GreetingClient: &clientstremmingpb.GreetingClientStream{FirstName: "Jane", LastName: "Smith"}},
		{GreetingClient: &clientstremmingpb.GreetingClientStream{FirstName: "Alice", LastName: "Johnson"}},
	}

	for _, req := range requests {
		log.Printf("Sending request: %v", req)
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending stream: %v", err)
		}
		time.Sleep(1 * time.Second) // Simulate some delay between requests
	}

	// Close the stream and receive the server's response
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	// Print the final response from the server
	log.Printf("Response from server: %v", res.GetResult())
}
