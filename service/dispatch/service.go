package dispatch

import (
	"context"

	pb "task/pb/dispatch"

	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedDispatchServiceServer
}

func (s *server) Ping(ctx context.Context, in *emptypb.Empty) (*pb.DispatchResponse, error) {
	return &pb.DispatchResponse{ErrorCode: 2000, ErrorMsg: "ok", Result: ""}, nil
}

func (s *server) Dispatch(ctx context.Context, req *pb.DispatchRequest) (*pb.DispatchResponse, error) {
	return &pb.DispatchResponse{ErrorCode: 2000, ErrorMsg: "ok", Result: ""}, nil
}

func NewDispatchService() *server {
	return &server{}
}
