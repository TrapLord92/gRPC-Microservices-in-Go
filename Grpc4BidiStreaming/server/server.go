package main

import (
	"io"
	"log"
	"net"

	"github.com/TrapLord92/gRPC-Microservices-in-Go.git/Grpc4BidiStreaming/bidipb"
	"google.golang.org/grpc"
)

// server is used to implement the BidiService
type server struct {
	bidipb.UnimplementedBidiServiceServer
}

// BidiChat implements the bidirectional streaming RPC
func (s *server) BidiChat(stream bidipb.BidiService_BidiChatServer) error {
	log.Println("BidiChat has been invoked.")

	// Continuously receive messages from the client and send responses
	for {
		// Receive a message from the client
		req, err := stream.Recv()
		if err == io.EOF {
			// Client has closed the stream
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
			return err
		}

		// Log the message received
		log.Printf("Received from client: %s", req.GetMessage())

		// Send a response back to the client
		err = stream.Send(&bidipb.BidiResponse{
			Reply: "Server says: " + req.GetMessage(),
		})
		if err != nil {
			log.Fatalf("Error sending response: %v", err)
			return err
		}
	}
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the BidiService with the server
	bidipb.RegisterBidiServiceServer(grpcServer, &server{})

	// Start the gRPC server
	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
