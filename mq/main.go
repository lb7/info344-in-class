package main

import (
	"github.com/streadway/amqp"
	"os"
	"fmt"
	"log"
)

func listen(msgs <-chan amqp.Delivery) {
	log.Println("Listening for new messages...")
	for msg := range msgs {
		log.Println(string(msg.Body))
	}
}

func main() {
	mqAddr := os.Getenv("MQADDR")
	if len(mqAddr) == 0 {
		mqAddr = "localhost:5672"
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s", mqAddr))
	if err != nil {
		log.Fatalf("Could not connect to message queue: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("error creating channel: %v", err)
	}

	q, err := channel.QueueDeclare("testQ", false, false, false, false, nil)

	msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)

	go listen(msgs)
	neverEnd := make(chan bool)
	<-neverEnd
}