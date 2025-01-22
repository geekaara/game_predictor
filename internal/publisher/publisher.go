package publisher

import "log"

func Publish(data []byte) error {
	log.Printf("Publishing data: %s\n", string(data))
	return nil
}
