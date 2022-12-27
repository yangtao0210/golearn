package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnErr(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErr(err, "Failed to open a channel")
	defer ch.Close()

	// 声明交换器
	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "Failed to declare an exchange")

	// 声明队列
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // 独占：连接关闭时，队列被自动删除
		false, // no-wait
		nil,   // arguments
	)
	failOnErr(err, "Failed to declare a queue")

	// 绑定队列-交换器
	ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)
	failOnErr(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("[x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
