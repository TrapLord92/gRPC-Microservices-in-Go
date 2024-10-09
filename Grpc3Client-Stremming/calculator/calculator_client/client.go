package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TrapLord92/gRPC-Microservices-in-Go.git/Grpc2Server-Stremming/calculator/calculatorpb"
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
	client := calculatorpb.NewCalculatorServiceClient(clientConnection)
	doUnary(client)
}

func doUnary(client calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 5,
	}
	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Println("The sum result is :", res.SumResult)
}
