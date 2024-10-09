package main

import (
	"context"
	"log"
	"net"

	"github.com/TrapLord92/gRPC-Microservices-in-Go.git/UnaryApi/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// Adjust the import path to where greetpb package is generated
)

// server is used to implement greet.GreetServiceServer
type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

// calculateSum implements the CalculateSum RPC method
func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("CalculateSum RPC invoked with: %v", req)
	first_number := req.FirstNumber
	second_number := req.SecondNumber
	sum := first_number + second_number
	result := &calculatorpb.SumResponse{SumResult: sum}
	return result, nil
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the GreetService server
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	// Enable reflection for gRPC server (optional but useful for testing with tools like grpcurl)
	reflection.Register(s)

	// Start the gRPC server
	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
