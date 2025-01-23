package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/geekaara/game_predictor/internal/fetcher"
	"github.com/geekaara/game_predictor/internal/publisher"
)

func main() {
	// 1. Read environment variables
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		log.Fatal("GCP_PROJECT_ID not set")
	}
	topicID := "mlb-live-feed" // adjust if your Pub/Sub topic has a different name

	// 2. Create a Pub/Sub client
	pubClient, err := publisher.NewPubSubClient(projectID, topicID)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	defer pubClient.Close()

	// 3. The specific MLB game we want to track
	//    Replace "716463" with the actual game PK you want.
	gamePK := "716463"

	// 4. Set up a ticker to fetch & publish every 30 seconds
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	log.Println("Starting MLB live data fetch loop. Press Ctrl-C to exit.")

	ctx := context.Background()

	for {
		select {
		case <-ticker.C:
			// Fetch the GUMBO data
			data, fetchErr := fetcher.FetchMLBData(ctx, gamePK)
			if fetchErr != nil {
				log.Printf("Error fetching MLB data: %v\n", fetchErr)
				continue
			}

			// Publish the data to Pub/Sub
			if err := pubClient.Publish(ctx, data); err != nil {
				log.Printf("Error publishing MLB data: %v\n", err)
			} else {
				log.Println("Successfully published MLB data to Pub/Sub")
			}
		}
	}
}
