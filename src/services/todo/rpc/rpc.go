package rpc

import (
	"context"
	"fmt"

	"todo.com/proto/greeting"
	pb "todo.com/proto/todo"
)

type GreeterServer struct {
	pb.GreeterServer
	gClient greeting.GreeterClient
}

func NewRPC(grpb greeting.GreeterClient) *GreeterServer {
	return &GreeterServer{
		gClient: grpb,
	}
}

func (g *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	resp, err := g.gClient.SayHello(ctx, &greeting.HelloRequest{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s from TODO Server. AND Get msg from GreetingServer \"%v\" Congrats! ", req.Name, resp.GetMessage()),
	}, nil
}

func (g *GreeterServer) SayHelloAgain(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s again from TODO Server", req.Name),
	}, nil
}
