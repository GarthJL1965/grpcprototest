package main

import (
	"log"
	"time"

	pb "github.com/gjlanc65/grpcprototest/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpc server address
const address = "localhost:8000"

func main() {
	log.Println("Quote Client ...")

	// Set up connection with the grpc server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	defer conn.Close()

	// Create a client instance
	c := pb.NewQuoteTransactionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	quote, err := c.MakeQuote(
		ctx,
		&pb.QuoteRequest{
			First:  -1,
			Second: -1,
			Third:  -1,
		},
	)
	if err != nil {
		log.Fatalf("Could not get quote: %v", err)
	}

	log.Printf("Quote: %s\n", quote.GetQuote())

}
