// Package api serves as client for grpc
// TODO: integrate this with main csv-reader logic. For now using as example
package api

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/syntax753/csver/cmd/csv-storer/api"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(fmt.Sprintf("%d", 7777), grpc.WithInsecure())
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
