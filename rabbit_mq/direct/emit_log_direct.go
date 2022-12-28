package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func dealError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func dealBody(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	dealError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	dealError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	dealError(err, "Failed to declare an exchange")

	body := dealBody(os.Args)
	err = ch.Publish(
		"logs_direct",
		severityFrom(os.Args),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	dealError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
