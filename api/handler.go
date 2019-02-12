// Package api stores the interface this project client server and is responsible
// for the proto definitons and build
package api

import (
	"fmt"

	context "golang.org/x/net/context"
)

// Server represents the CSV storer grpc API
type Server struct{}

// Store accepts a grpc request with the csv payload line and returns sync response for now
func (s *Server) Store(ctx context.Context, in *StoreRequest) (*StoreResponse, error) {
	fmt.Printf("Received message %s\n", in.GetRequest())
	return &StoreResponse{Response: "Received"}, nil
}
