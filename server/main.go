package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/gurasissingh/grpc-pprof/sleep"
	"github.com/gurasissingh/grpc-pprof/json"
	pb "github.com/gurasissingh/grpc-pprof/protobuf"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

// server is used to implement protobuf.GreeterServer.
type server struct {
}

func (s *server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	json.JsonMarshal()
	sleep.SleepFunc()
	return &pb.HelloResponse{Message: "Returning Input:  " + in.GetInput()}, nil
}

func main() {

	runtime.SetBlockProfileRate(1)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	go func() {
		http.ListenAndServe(":8024", nil)
	}()

	server := server{}
	// setup server
	srv := grpc.NewServer()
	pb.RegisterGrpcServiceServer(srv, &server)
	fmt.Printf("Server Started")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
