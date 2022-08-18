package dispatch

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client DispatchServiceClient
)

func init() {
	conn, err := grpc.Dial("task_dispatch:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = NewDispatchServiceClient(conn)
	fmt.Println("init dispatch client")
}

func Ping(ctx context.Context) error {
	resp, err := client.Ping(ctx, &Empty{})
	if err != nil {
		log.Fatalf("ping error: %v", err)
		return err
	}
	log.Printf("ping result: %+v", resp)
	return nil
}

func Dispatch(ctx context.Context, event string, resourceId int64, dagId, processorId int) ([]int64, error) {
	req := &DispatchRequest{
		Event:       event,
		ResourceId:  resourceId,
		DagId:       int64(dagId),
		ProcessorId: int64(processorId),
	}
	resp, err := client.Dispatch(ctx, req)
	if err != nil {
		log.Fatalf("dispatch error: %v", err)
		return nil, err
	}
	log.Printf("dispatch result: %s", resp)
	return resp.ProcessorIdList, nil
}
