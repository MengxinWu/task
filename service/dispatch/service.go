package dispatch

import (
	"context"
	"fmt"
	"task/models"
	pb "task/pb/dispatch"

	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedDispatchServiceServer
}

func (s *server) Ping(ctx context.Context, in *emptypb.Empty) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "ok"}, nil
}

func (s *server) Dispatch(ctx context.Context, req *pb.DispatchRequest) (*pb.DispatchResponse, error) {
	var (
		hdl DispatchHandler
		ok  bool
		err error
	)
	// 调度算法
	dispatchEvent := &models.DispatchEvent{
		Event:       req.Event,
		ResourceId:  req.ResourceId,
		DagId:       int(req.DagId),
		ProcessorId: int(req.ProcessorId),
	}

	if hdl, ok = EventHandlerMap[req.Event]; !ok {
		return nil, fmt.Errorf("event not exist: %s", req.Event)
	}

	if err = hdl.Prepare(ctx, dispatchEvent); err != nil {
		return nil, err
	}

	if err = hdl.Compute(ctx, dispatchEvent); err != nil {
		return nil, err
	}

	if err = hdl.After(ctx, dispatchEvent); err != nil {
		return nil, err
	}

	return &pb.DispatchResponse{
		ProcessorIdList: dispatchEvent.ExecutorList,
	}, nil
}

func NewDispatchService() *server {
	return &server{}
}
