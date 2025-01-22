package main

import (
	"log"
	"os"
	"time"

	"github.com/geekaara/game_predictor/internal/fetcher"
	"github.com/geekaara/game_predictor/internal/publisher"
)

func main() {
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		log.Fatal("GCP_PROJECT_ID not set")
	}

	// TODO: create PubSub client, ticker, fetch logic, etc.
	log.Println("Project ID:", projectID)

	// This is just an example to verify your code compiles
	time.Sleep(3 * time.Second)
	log.Println("Done.")
}
