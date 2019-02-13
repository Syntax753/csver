// Package handler is responsible for providing the busines logic to the inbound request to the server
package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/syntax753/csver/api"

	"github.com/syntax753/csver/lib/model"
)

// Server represents the CSV storer grpc API
type Server struct {
}

var (
	database = make(dbType)
)

// dbtype represents the db conceptual persistence layer
type dbType map[string]model.DataRecord

// Store accepts a grpc request with the csv payload line and returns sync response for now
func (s *Server) Store(ctx context.Context, in *api.StoreRequest) (*api.StoreResponse, error) {
	req := in.GetRequest()
	// log.Printf("Processing %s\n", req)

	if req == "" {
		log.Printf("Database: %+v", database)

		os.Exit(0)
	}

	var rec = model.Record{}
	err := json.Unmarshal([]byte(req), &rec)

	if err != nil {
		return &api.StoreResponse{Response: "Error"}, err
	}

	sts := "Added"
	_, ok := database[rec.ID]

	if ok {
		sts = "Updated"
	}

	database[rec.ID] = rec.Record

	return &api.StoreResponse{Response: fmt.Sprintf("%v: %v\n", rec.ID, sts)}, nil
}
