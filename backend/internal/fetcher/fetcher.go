package fetcher

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// FetchMLBData calls the MLB Stats API for a specific game PK (GUMBO feed)
// and returns the full JSON response as a byte slice.
func FetchMLBData(ctx context.Context, gamePK string) ([]byte, error) {
	url := fmt.Sprintf("https://statsapi.mlb.com/api/v1.1/game/%s/feed/live", gamePK)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d from MLB API", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read MLB API response: %w", err)
	}

	return data, nil
}
