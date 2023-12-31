package main

import(
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main(){
	// Connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	printErrorAndExit(err, "Failed to open a channel")
	defer conn.Close()

	// Channel
	ch, err := conn.Channel()
	printErrorAndExit(err, "Failed to open a channel")
	defer ch.Close()

	// Exchange
	err = ch.ExchangeDeclare(
		"directExchange", // name
		"direct", //type
		false, // durable
		true, // auto deleted
		false, // internal
		false, // no-wait
		nil, //arguements
	)
	printErrorAndExit(err, "Failed to declare an exchange")

	// Declare and bind queue
	q, err := ch.QueueDeclare(
		"", // name, empty, let server generate id
		false, // durable
		true, // delete when unused
		true, // exclusive
		false, // no-wait
		nil, // arguments 
	)
	printErrorAndExit(err, "Failed to declare a queue")
	err = ch.QueueBind(
		q.Name, // queue name
		"one", // routing key
		"directExchange", // exchange
		false,
		nil,
	)
	printErrorAndExit(err, "Failed to bind a queue")

	// Consume
	msgs, err := ch.Consume(
		q.Name, // queue
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	printErrorAndExit(err, "Failed to register a consumer")

	go func(){
		for d := range msgs{
			bodyString := string(d.Body)
			fmt.Println("Received: ", bodyString)
			d.Ack(false)
		}
	}()

	fmt.Println("Waiting for msgs")
	forever := make(chan bool)
	<-forever
}

func printErrorAndExit(err error , msg string) { 
	if err != nil { 
		log.Fatalln(msg, ":", err)
	}
}