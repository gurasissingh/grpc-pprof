package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/gurasissingh/grpc-pprof/protobuf"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcServiceClient(conn)

	for i := 0; i < 100; i++ {
		// Contact the server and print out its response.
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		r, err := c.Hello(ctx, &pb.HelloRequest{Input: "test data"})
		if err != nil {
			fmt.Printf("could not get data: %v", err)
		}
		fmt.Printf("%s\n", r.GetMessage())
	}
}
