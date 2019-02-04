package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/syntax753/csver/lib/model"

	"github.com/tkanos/gonfig"
)

var (
	reader *csv.Reader
)

func init() {
	fmt.Println("Init reader")
	cfg := Config{}

	err := gonfig.GetConf("./dev.json", &cfg)
	if err != nil {
		log.Fatalf("Could not load configuration! %v\n", err)
	}
	fmt.Printf("Using config: %+v\n", cfg)

	f, err := os.Open(cfg.Data)
	if err != nil {
		log.Fatalf("Could not open data file %q! %v\n", cfg.Data, err)
	}

	reader = csv.NewReader(f)
	reader.Comma = []rune(cfg.Separator)[0]
	fmt.Println("Init done")
}

// Config holds the configuration for reader and the endpoint to store the data
type Config struct {
	Database struct {
		Host string
		Port string
	}
	Data      string
	Separator string
}

func main() {

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Error importing %v - igoring\n", line)
			continue
		}

		data := &model.DataRecord{
			Name:        line[1],
			Email:       line[2],
			CountryCode: "+44",
			Mobile:      line[3],
		}

		record := &model.Record{
			ID:     line[0],
			Record: *data,
		}

		fmt.Printf("Processing %+v\n", record)

	}

}
