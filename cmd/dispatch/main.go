package main

import (
	"fmt"
	"net"

	"task/driver"
	pb "task/pb/dispatch"
	"task/service/dispatch"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// 初始化资源
	driver.InitEngine()
	driver.InitExecuteConn()

	// 初始化调度管理器
	service.InitDispatchHandler()

	// 创建并运行grpc服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":50051"))
	if err != nil {
		log.Panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterDispatchServiceServer(s, service.NewDispatchService())
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Panic(err)
	}
}
