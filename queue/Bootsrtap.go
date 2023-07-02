package queue

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var ChannelRabbitMQ *amqp.Channel

func Connect() {
	Port := os.Getenv("AMQP_PORT")
	Password := os.Getenv("AMQP_PASSWORD")
	Username := os.Getenv("AMQP_USER")
	amqpServerURL := "amqp://" + Username + ":" + Password + "@localhost:" + Port + "/"

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect Successful to RabbitMQ")

	defer connectRabbitMQ.Close()

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	ChannelRabbitMQ, err = connectRabbitMQ.Channel()
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}

	fmt.Println("Connect Successful to RabbitMQ Channel")

	defer ChannelRabbitMQ.Close()

	if err != nil {
		failOnError(err, "Failed to Declare RabbitMQ Queue")
	}

	failOnError(err, "Success Declare RabbitMQ Queue")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func testPub() {
	q, err := ChannelRabbitMQ.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ChannelRabbitMQ.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
