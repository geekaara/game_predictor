package main

import (
	"log"
	"os"

	// IMPORTANT: matches your module path above, plus /internal/...
	"github.com/geekaara/game_predictor/internal/fetcher"
	"github.com/geekaara/game_predictor/internal/publisher"
)

func main() {
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		log.Fatal("GCP_PROJECT_ID not set")
	}

	// Just a test: fetch some data, then publish it.
	data, err := fetcher.FetchMLBData("123456")
	if err != nil {
		log.Fatal("error fetching MLB data:", err)
	}
	err = publisher.Publish(data)
	if err != nil {
		log.Fatal("error publishing data:", err)
	}

	log.Println("All done!")
}
