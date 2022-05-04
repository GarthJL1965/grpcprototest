package main

import (
	"fmt"

	pb "github.com/gjlanc65/grpcprototest/protofiles"
	"google.golang.org/protobuf/proto"
)

func main() {

	// using the proto created struct
	q := &pb.QuoteRequest{
		First:  -1,
		Second: -1,
		Third:  -1,
	}

	// Serializing the struct and assigning it to body
	body, _ := proto.Marshal(q)

	// De-serializing the body and saving it to p1 for testing
	q1 := &pb.QuoteRequest{}
	_ = proto.Unmarshal(body, q1)

	fmt.Println("Original struct loaded from proto file:", q)
	fmt.Println("Marshalled proto data: ", body)
	fmt.Println("Unmarshalled struct: ", q1)

}
