// Package handler is responsible for providing the busines logic to the inbound request to the server
package handler

import (
	"fmt"

	"github.com/syntax753/csver/api"

	context "golang.org/x/net/context"
)

// Server represents the CSV storer grpc API
type Server struct{}

// Store accepts a grpc request with the csv payload line and returns sync response for now
func (s *Server) Store(ctx context.Context, in *api.StoreRequest) (*api.StoreResponse, error) {
	fmt.Printf("Received message %s\n", in.GetRequest())
	return &api.StoreResponse{Response: "Received"}, nil
}
