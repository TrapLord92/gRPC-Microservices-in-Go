package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/TrapLord92/gRPC-Microservices-in-Go.git/Grpc3Client-Stremming/greet/greetpb/clientstremmingpb"
	"google.golang.org/grpc"
)

// server is used to implement greetclientstremming.GreetClientStreamServiceServer
type server struct {
	clientstremmingpb.UnimplementedGreetClientStreamServiceServer
}

// GreetStreeming implements the client-side streaming RPC
func (s *server) GreetStreeming(stream clientstremmingpb.GreetClientStreamService_GreetStreemingServer) error {
	fmt.Println("GreetStreeming invoked with a client-streaming request")

	var result string

	// Receive the stream of requests from the client
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Once the client finishes sending the stream, send the final response
			return stream.SendAndClose(&clientstremmingpb.GreetClientStreamResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while receiving stream: %v", err)
			return err
		}

		// Process the incoming stream request
		firstName := req.GetGreetingClient().GetFirstName()
		lastName := req.GetGreetingClient().GetLastName()
		result += "Hello, " + firstName + " " + lastName + "! "
	}
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the GreetClientStreamService server
	clientstremmingpb.RegisterGreetClientStreamServiceServer(s, &server{})

	// Start the gRPC server
	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
