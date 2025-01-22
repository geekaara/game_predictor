package fetcher

import "fmt"

// FetchMLBData is a stub function for now
func FetchMLBData(gamePK string) ([]byte, error) {
	return []byte(fmt.Sprintf("Fake fetch data for game %s", gamePK)), nil
}
