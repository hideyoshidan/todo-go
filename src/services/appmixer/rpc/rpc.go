package rpc

import (
	"context"

	"todo.com/proto/appmixer"
)

type AppMixerServer struct {
	appmixer.AppmixerServer
}

func NewRPC() *AppMixerServer {
	return &AppMixerServer{}
}

func (g *AppMixerServer) SayHello(ctx context.Context, req *appmixer.AppRequest) (*appmixer.AppResponse, error) {
	return &appmixer.AppResponse{
		Message: "TEST",
	}, nil
}

// Sends another greeting
func (g *AppMixerServer) SayHelloAgain(ctx context.Context, req *appmixer.AppRequest) (*appmixer.AppResponse, error) {
	return &appmixer.AppResponse{
		Message: "TEST",
	}, nil
}
