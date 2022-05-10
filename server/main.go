package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/gjlanc65/grpcprototest/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/reflection"
)

const (
	defaultPort = ":8000"
)

type server struct {
	pb.UnimplementedQuoteTransactionServer
}

func main() {
	log.Println("Quote Server starting up ...")

	// NewServer creates a gRPC server which has no service registered and has not started
	// to accept requests yet.
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", defaultPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// We are making use of the function that compiled proto made for us to register
	// our GRPC server so that the clients can make use of the functions tied to our
	// server remotely via the GRPC server (like MakeQuote function)

	// The first argument is the grpc server instance
	// The second argument is the service who's methods we want to expose (in our case)
	// we have put it in this program only
	pb.RegisterQuoteTransactionServer(s, &server{})

	// Serve accepts incoming connections on the listener lis, creating a new ServerTransport
	// and service goroutine for each. The service goroutines read gRPC requests and then
	// call the registered handlers to reply to them.

	go func() {
		log.Println("Server running on tcp (localhost):8000")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	log.Println("Stopping the server")
	s.Stop()
	log.Println("Closing the listener")
	lis.Close()
	log.Println("Server has been Shut down")

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
	// Returning a response of type QuoteResponse
	return &pb.QuoteResponse{Quote: "willy shakeit"}, nil
}
