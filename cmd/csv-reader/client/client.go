// Package api serves as client for grpc
// TODO: integrate this with main csv-reader logic. For now using as example
package client

import (
	"fmt"
	"log"

	"github.com/syntax753/csver/api"
	"github.com/syntax753/csver/lib/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func init() {
	
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(fmt.Sprintf("%d", config.Global().GRPCPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client could not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewStorerClient(conn)

	r, err := c.Store(context.Background(), &api.StoreRequest{
		Request: "Test",
	})

	if err != nil {
		log.Fatalf("Client could not connect: %s", err)
	}
	log.Printf("Response from server: %s", r.Response)

}
