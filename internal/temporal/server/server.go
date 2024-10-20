package server

import (
	pb "fx-web/internal/temporal/pb"
	"google.golang.org/grpc"
	"net"
)

type TemporalServer struct {
	grpcServer *grpc.Server
	// 其他必要的字段
}

// TODO 在实现这个函数时，存在包等级的问题
// 需要进一步理解golang的包管理机制和可见性问题
func NewTemporalServer(workflowService *pb.WorkflowService) *TemporalServer {
	grpcServer := grpc.NewServer()
	pb.RegisterWorkflowServiceServer(grpcServer, workflowService)
	return &TemporalServer{grpcServer: grpcServer}
}

func StartTemporalServer(server *TemporalServer) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	return server.grpcServer.Serve(lis)
}
