package dispatch

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client DispatchServiceClient
)

func init() {
	conn, err := grpc.Dial("dispatch:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = NewDispatchServiceClient(conn)
}

func Ping() {
	r, err := client.Ping(context.Background(), &Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("result: %s", r)
}
