package main

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

func init() {
	fmt.Println("Initialising reader")
	cfg := Config{}

	// wd, _ := os.Getwd()
	err := gonfig.GetConf("./dev.json", &cfg)
	if err != nil {
		fmt.Printf("Could not load configuration! %v\n", err)
		panic("Aborting...")
	}

	fmt.Printf("Using config: %+v\n", cfg)
}

// Config holds the configuration for reader and the endpoint to store the data
type Config struct {
	Database struct {
		Host string
		Port string
	}
	Separator string
}


func main() {

}
