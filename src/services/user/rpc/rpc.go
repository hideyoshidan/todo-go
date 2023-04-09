package rpc

import (
	"context"
	"fmt"

	pb "todo.com/proto/user"
)

type GreeterServer struct {
	pb.GreeterServer
}

func NewRPC() *GreeterServer {
	return &GreeterServer{}
}

func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s From User Server!", req.Name),
	}, nil
}

func (s *GreeterServer) SayHelloAgain(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s Again From User Server!", req.Name),
	}, nil
}
