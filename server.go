package main

import (
	"context"
	"log"
	"time"

	pb "ls/"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequests) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetExample())
	return &pb.HelloResponse{Message: "Hello, " + in.GetExample()}, nil
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	name := "World"
	res, err := c.SayHello(ctx, &pb.HelloRequest{Example: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", res.GetMessage())
}
