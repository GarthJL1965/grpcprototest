# grpcprototest
Go(lang) grpc &amp;&amp; protobuf explore - writing a service that returns the 3 segments required for a Shakespearean quote 

Sending -1 for a segment (first|second|third):
```protobuf
message QuoteRequest {
  int32       first   = 1;
  int32       second  = 2;
  int32       third   = 3;
}
```
says 'use a random number [0,#size-1]' for that segment

## (Windows) setup
* install protoc https://developers.google.com/protocol-buffers/docs/downloads#release-packages and add to path (manual)
* Might need ```go get google.golang./protobuf/cmd/protoc-gen-go@latest``` before install step
* ```go install google.golang.org/protobuf/cmd/protoc-gen-go@latest```
* Might need ```go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest```
* ```go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest```
* (generate proto files) - ```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protofiles/shakespeare.proto```
* ```go get google.golang.org/grpc```
* (server) imports "google.golang.org/grpc" && 	_ "google.golang.org/grpc/reflection"
* (client) imports "google.golang.org/grpc"
* (server && client) imports pb "github.com/gjlanc65/grpcprototest/protofiles"

## Internal notes
* server 'traps' Control-C/Break and terminates
* client is 'one shot'