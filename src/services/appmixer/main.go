package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "todo.com/proto/appmixer"
	"todo.com/services/appmixer/rpc"
)

const (
	RPC_PORT     = 4002
	GATEWAY_PORT = 8000
)

func main() {
	log.Print("Appmixer gRPC Server Starting")
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_PORT))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	rpc := rpc.NewRPC()
	// Attach the Greeter service to the server
	pb.RegisterAppmixerServer(s, rpc)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%d", RPC_PORT),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	log.Print("Appmixer grpc server waiting proxy...")

	mux := runtime.NewServeMux()
	err = pb.RegisterAppmixerHandler(context.Background(), mux, conn)

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", GATEWAY_PORT),
		Handler: mux,
	}

	log.Print("Appmixer grpc-gateway waiting request...")

	log.Fatalln(gwServer.ListenAndServe())
}
