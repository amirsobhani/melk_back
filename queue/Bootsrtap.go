package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var ChannelRabbitMQ *amqp.Channel

type Config struct {
	DefaultQueue     string
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	TimeOut          time.Duration
}

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

	//defer connectRabbitMQ.Close()

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	ChannelRabbitMQ, err = connectRabbitMQ.Channel()
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}

	fmt.Println("Connect Successful to RabbitMQ Channel")

	//defer ChannelRabbitMQ.Close()

	if err != nil {
		failOnError(err, "Failed to Declare RabbitMQ Queue")
	}

	fmt.Println("Success Declare RabbitMQ Queue")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (config *Config) PublishToQueue(v interface{}) {
	config.defaultSetter()

	q, err := ChannelRabbitMQ.QueueDeclare(
		config.DefaultQueue,     // name
		config.Durable,          // durable
		config.DeleteWhenUnused, // delete when unused
		config.Exclusive,        // exclusive
		config.NoWait,           // no-wait
		nil,                     // arguments
	)

	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), config.TimeOut)

	defer cancel()

	body, err := json.Marshal(v)

	failOnError(err, "Failed to Marshal value")

	err = ChannelRabbitMQ.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	failOnError(err, "Failed to publish a message")
}

func (config *Config) defaultSetter() {
	if config.DefaultQueue == "" {
		config.DefaultQueue = os.Getenv("AMQP_DEFAULT_QUEUE")
	}

	if &config.Durable == nil {
		config.Durable = false
	}

	if &config.Exclusive == nil {
		config.Exclusive = false
	}

	if &config.NoWait == nil {
		config.NoWait = false
	}

	if &config.DeleteWhenUnused == nil {
		config.DeleteWhenUnused = false
	}

	if config.TimeOut == 0 {
		config.TimeOut = 5 * time.Second
	}
}
