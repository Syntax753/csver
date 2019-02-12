package main

import (
	"fmt"
	"log"
	"net"

	"github.com/syntax753/csver/lib/config"

	"github.com/syntax753/csver/api"
	"google.golang.org/grpc"
)

func init() {
	log.Println("Init server")
	log.Println("Server done")
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Global().GRPCPort))
	if err != nil {
		log.Fatalf("grpc listener failed: %v", err)
	}

	// create a server instance
	s := api.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterStorerServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc listener failed: %v", err)
	}
}
