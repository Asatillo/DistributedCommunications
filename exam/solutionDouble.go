package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

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

	// Declare and bind queue
	q, err := ch.QueueDeclare(
		"Double", 
		false,    
		true,     
		true,     
		false,    
		nil,      
	)
	printErrorAndExit(err, "Failed to declare a queue")
	err = ch.QueueBind(
		q.Name,             
		"anykey3",          
		"p3fanoutExchange", 
		false,
		nil,
	)
	printErrorAndExit(err, "Failed to bind a queue")

	// Consume
	msgs, err := ch.Consume(
		q.Name, 
		"double",
		false,
		false,
		false,
		false,
		nil,
	)
	printErrorAndExit(err, "Failed to register a consumer")

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for d := range msgs {
			if string(d.Body) != "END" {
				num, _ := strconv.Atoi(string(d.Body))
				fmt.Println("DOUBLE Received: ", num, " Result: ", num*2)
			}else{
				fmt.Println("Got END message")
				ch.Close()
			}
			d.Ack(false)
		}
	}()
	wg.Wait()
	fmt.Println("Finish")

	fmt.Println("Waiting for msgs")
	forever := make(chan bool)
	<-forever
}

func printErrorAndExit(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, ":", err)
	}
}
