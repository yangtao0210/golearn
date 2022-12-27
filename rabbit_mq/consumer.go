package rabbit_mq

import (
	"github.com/streadway/amqp"
	"log"
)

func Receive(qName string) {
	// 1.连接 rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to rabbitmq")
	defer conn.Close()
	// 2. 创建通道channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	// 3.声明要接收的消息队列
	q, err := ch.QueueDeclare(
		qName,
		false,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to declare a queue")
	// 4.获取接收消息的Delivery通道
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message : %s", string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
