package main

import (
	"context"
	"fmt"
	"grpc-helloworld/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Saurabh"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	fmt.Printf("Greeting: %s\n", r.Message)
}
