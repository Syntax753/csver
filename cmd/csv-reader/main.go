package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	csver "github.com/syntax753/csver/cmd/csv-reader/client"

	"github.com/syntax753/csver/lib/config"
)

var (
	reader *csv.Reader
	client *csver.Client
)

func init() {
	log.Println("Init reader")

	f, err := os.Open(config.Global().Data)
	if err != nil {
		log.Fatalf("Could not open data file %q! %v\n", config.Global().Data, err)
	}

	reader = csv.NewReader(f)
	reader.Comma = []rune(config.Global().Separator)[0]

	client = csver.NewClient()

	log.Println("Reader done")
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

		// data := &model.DataRecord{
		// 	Name:        line[1],
		// 	Email:       line[2],
		// 	CountryCode: "+44",
		// 	Mobile:      line[3],
		// }

		// record := &model.Record{
		// 	ID:     line[0],
		// 	Record: *data,
		// }

		// log.Println("Processing %+v\n", record)

		client.Process("Hello")

	}
}
