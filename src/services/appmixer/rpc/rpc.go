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

func (s *AppMixerServer) SayHello(ctx context.Context, req *appmixer.AppRequest) (*appmixer.AppResponse, error) {
	return &appmixer.AppResponse{
		Message: "TEST",
	}, nil
}

// Sends another greeting
func (s *AppMixerServer) SayHelloAgain(ctx context.Context, req *appmixer.AppRequest) (*appmixer.AppResponse, error) {
	return &appmixer.AppResponse{
		Message: "TEST",
	}, nil
}

// Sign Up
func (s *AppMixerServer) SignUp(ctx context.Context, req *appmixer.SignUpRequest) (*appmixer.SignUpResponse, error) {
	return &appmixer.SignUpResponse{
		Token: "sample",
	}, nil
}

// Sign In
func (s *AppMixerServer) SignIn(ctx context.Context, req *appmixer.SignInRequest) (*appmixer.SignInResponse, error) {
	return &appmixer.SignInResponse{
		Token: "sample",
	}, nil
}
