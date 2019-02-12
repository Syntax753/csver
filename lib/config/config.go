// Package config handles configuration applicable to the entire omnirepo
// Uses gonfig to
package config

import (
	"fmt"
	"log"

	"github.com/tkanos/gonfig"
)

var (
	cfg Configuration
)

func init() {
	log.Println("Init config")

	// //TODO: env variable picking the config values for the right env
	// path, err := filepath.Abs("./dev.json")

	// if err != nil {
	// 	log.Fatalf("Could not load configuration! %v\n", err)
	// }

	cfg = Configuration{}

	err := gonfig.GetConf("../../lib/config/dev.json", &cfg)
	if err != nil {
		log.Fatalf("Could not load configuration! %v\n", err)
	}

	fmt.Printf("%+v\n", cfg)

	log.Println("Config done")
}

// Global holds the configuration for the monorepo
func Global() Configuration {
	return cfg
}

// Configuration holds the configuration for reader and the endpoint to store the data
type Configuration struct {
	Database struct {
		Host string
		Port string
	}
	Data      string
	Separator string
	GRPCPort  int
}
