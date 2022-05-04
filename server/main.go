package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/gjlanc65/grpcprototest/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedQuoteTransactionServer
}

func main() {
	fmt.Println("Quote Server ...")

	// NewServer creates a gRPC server which has no service registered and has not started
	// to accept requests yet.
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// We are making use of the function that compiled proto made for us to register
	// our GRPC server so that the clients can make use of the functions tide to our
	// server remotely via the GRPC server (like MakeTransaction function)

	// The first argument is the grpc server instance
	// The second argument is the service who's methods we want to expose (in our case)
	// we have put it in this program only
	pb.RegisterQuoteTransactionServer(s, &server{})

	// Serve accepts incoming connections on the listener lis, creating a new ServerTransport
	// and service goroutine for each. The service goroutines read gRPC requests and then
	// call the registered handlers to reply to them.
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

// [ctx] is used by the goroutines to interact with GRPC
// [in] is the type of TransactionRequest
/*
   This function signature matches the service that we mentioned in the protobuf
*/
func (s *server) MakeQuote(ctx context.Context, in *pb.QuoteRequest) (*pb.QuoteResponse, error) {
	// Business logic will come here
	fmt.Println("Got first  ", in.First)
	fmt.Println("Got second ", in.Second)
	fmt.Println("Got third  ", in.Third)
	// Returning a response of type Transaction Response
	return &pb.QuoteResponse{Quote: "willy shakeit"}, nil
}
