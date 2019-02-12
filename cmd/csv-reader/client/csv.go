// Package csv serves includes all csv client
package csver

import (
	"context"
	"fmt"
	"log"

	"github.com/syntax753/csver/api"
	"github.com/syntax753/csver/lib/config"
	"google.golang.org/grpc"
)

var (
	storerClient api.StorerClient
)

// Client is the client which handles reading csv files and delegates the calls to
// the csv backend storage server
type Client struct {
	protoClient api.StorerClient
}

// TODO: pool of connections? maintaining connection perhaps?
// Currently opening new connection in #Process method on each call
// 
func init() {
	log.Println("Init client")

	log.Println("Client connection checked")
}

// NewClient returns the client responsible for parsing csv and communicating over the grpc api
func NewClient() *Client {
	return &Client{
		protoClient: storerClient,
	}
}

// Process delegates the handling of one string overo grpc
func (client *Client) Process(req string) {

	// // TODO: connect to server at construction time/rather than initially
	// Presumes one instance here for the purposes of this challenge
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", config.Global().GRPCPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client could not connect: %s", err)
	}
	defer conn.Close()
	storerClient = api.NewStorerClient(conn)

	r, err := storerClient.Store(context.Background(), &api.StoreRequest{
		Request: "Test",
	})

	if err != nil {
		log.Fatalf("Client could not connect: %s", err)
	}

	log.Printf("Response from server: %s", r.Response)
}
