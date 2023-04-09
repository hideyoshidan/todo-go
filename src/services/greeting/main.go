package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "todo.com/proto/greeting"
	rpc "todo.com/services/greeting/rpc"
)

const (
	RPC_PORT = 4000
)

func main() {
	log.Print("Greeting Server Starting")

	listenProt, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc := grpc.NewServer()
	rpc := rpc.NewRPC()
	pb.RegisterGreeterServer(grpc, rpc)
	reflection.Register(grpc)

	if err := grpc.Serve(listenProt); err != nil {
		log.Fatal(err)
	}
}
