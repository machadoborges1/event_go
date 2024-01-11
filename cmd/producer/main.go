package main

import rabbitmq "github.com/machadoborges1/event_go/pkg/rabbitmg"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World!", "amq.direct")
}
