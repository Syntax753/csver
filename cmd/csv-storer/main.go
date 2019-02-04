package main

import (
	"fmt"
	"log"
	"net"

	"github.com/syntax753/csver/cmd/csv-storer/api"
	"google.golang.org/grpc"
)

func main() {
	//TODO: Add port to config for server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
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
