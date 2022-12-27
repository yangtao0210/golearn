package rabbit_mq

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

func Send(qName, msg string) {
	// 1.连接 rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to rabbitmq")
	defer conn.Close()
	// 2. 创建通道channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	// 3.声明要发送到的队列
	q, err := ch.QueueDeclare(
		qName,
		false,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to declare a queue")
	// 4.将消息发布到声明的队列
	body := msg
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
