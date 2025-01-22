package publisher

import "fmt"

// Publish is a stub function for now
func Publish(data []byte) error {
	fmt.Println("Publishing data:", string(data))
	return nil
}
