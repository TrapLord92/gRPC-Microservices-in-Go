package main

import (
	"fmt"
	"log"

	"github.com/TrapLord92/gRPC-Microservices-in-Go.git/Basic/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Hello, gRPC!")

	// Set up a context with a timeout for the connection attempt.

	// Use grpc.DialContext instead of grpc.Dial (handles cancellation, timeouts)
	clientConnection, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // Insecure credentials for non-TLS
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer clientConnection.Close()

	// Create a new GreetService client using the established connection
	client := greetpb.NewGreetServiceClient(clientConnection)
	fmt.Printf("Client created successfully! Client: %v\n", client)
}
