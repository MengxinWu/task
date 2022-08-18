package dispatch

import (
	"context"
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
	// todo 调度算法
	return &pb.DispatchResponse{
		ProcessorIdList: []int64{1},
	}, nil
}

func NewDispatchService() *server {
	return &server{}
}
