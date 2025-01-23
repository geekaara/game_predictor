package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
)

// PubSubClient wraps a Pub/Sub client and a topic reference
type PubSubClient struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

// NewPubSubClient initializes the Pub/Sub client and gets the topic reference
func NewPubSubClient(projectID, topicID string) (*PubSubClient, error) {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %w", err)
	}

	// Make sure the topic exists. If it doesn't, you can optionally create it here.
	topic := client.Topic(topicID)
	return &PubSubClient{
		client: client,
		topic:  topic,
	}, nil
}

// Publish sends the data to Pub/Sub as a single message
func (p *PubSubClient) Publish(ctx context.Context, data []byte) error {
	result := p.topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish: %w", err)
	}
	log.Printf("Published message with ID: %s", id)
	return nil
}

// Close gracefully shuts down the Pub/Sub client connection
func (p *PubSubClient) Close() {
	if err := p.client.Close(); err != nil {
		log.Printf("Error closing PubSub client: %v", err)
	}
}
