package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

func main() {
	// Connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	printErrorAndExit(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Channel
	ch, err := conn.Channel()
	printErrorAndExit(err, "Failed to open a channel")
	defer ch.Close()

	// Exchange
	err = ch.ExchangeDeclare(
		"p3fanoutExchange",
		"fanout",
		false,
		true,
		false,
		false,
		nil,
	)
	printErrorAndExit(err, "Failed to declare an exchange")

	//Publish Messages
	for i := 0; i <= 10; i++ {
		publishMsg(ch, "p3fanoutExchange", "anykey1", strconv.Itoa(i))
	}
	publishMsg(ch, "p3fanoutExchange", "anykey1", "END")
	fmt.Println("END...Finished sending numbers!")
}

func printErrorAndExit(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, ":", err)
	}
}

func publishMsg(c *amqp.Channel, ex string, key string, msg string) {
	body := msg
	err := (*c).Publish(
		ex,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	printErrorAndExit(err, "Failed to publish a message")
	fmt.Println("Sent: ", body)
}
