package dispatch

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	GrpcAddress = "task_dispatch:50051"
)

var (
	client DispatchServiceClient
)

func init() {
	conn, err := grpc.Dial(GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	client = NewDispatchServiceClient(conn)
	log.Printf("init dispatch grpc client success...")
	return
}

// Ping ping.
func Ping(ctx context.Context) error {
	resp, err := client.Ping(ctx, &Empty{})
	if err != nil {
		log.Errorf("ping rpc error: %v", err)
		return err
	}
	log.Printf("ping rpc result: %+v", resp)
	return nil
}

// Dispatch dispatch.
func Dispatch(ctx context.Context, event string, resourceId int64, dagId, processorId int) ([]int64, error) {
	req := &DispatchRequest{
		Event:       event,
		ResourceId:  resourceId,
		DagId:       int64(dagId),
		ProcessorId: int64(processorId),
	}
	resp, err := client.Dispatch(ctx, req)
	if err != nil {
		log.Errorf("dispatch rpc error: %v", err)
		return nil, err
	}
	log.Printf("dispatch rpc result: %s", resp)
	return resp.ProcessorIdList, nil
}
