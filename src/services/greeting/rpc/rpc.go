package rpc

import (
	"context"
	"fmt"

	pb "todo.com/proto/greeting"
)

type GreeterServer struct {
	pb.GreeterServer
}

// new greeting server
func NewRPC() *GreeterServer {
	return &GreeterServer{}
}

func (g *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}

func (g *GreeterServer) SayHelloAgain(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s again!", req.Name),
	}, nil
}
