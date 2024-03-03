package main

import (
	"fmt"
	"time"

	"github.com/rielj/go-message-broker/internal/api"
)

func main() {
	broker := api.NewBroker()

	subscriber := broker.Subscribe("example_topic")
	go func() {
		for {
			select {
			case message, ok := <-subscriber.Channel:
				if !ok {
					fmt.Println("Subscriber Channel closed")
					return
				}
				fmt.Printf("Received message: %v\n", message)
			case <-subscriber.Unsubscribe:
				fmt.Println("Unsubscribed")
				return
			}
		}
	}()

	broker.Publish("example_topic", "Hello, World!")
	broker.Publish("example_topic", "This is a message")

	time.Sleep(2 * time.Second)
	broker.Unsubscribe("example_topic", subscriber)

	broker.Publish("example_topic", "This message should not be received")

	time.Sleep(2 * time.Second)
}
