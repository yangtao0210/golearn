package rabbit_mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func main() {
	// 1.连接rabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("connect to RabbitMQ failed, err:%v\n", err)
		return
	}
	defer conn.Close()

	// 2. 通过连接 建立通道channel
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("open a channel failed, err:%v\n", err)
		return
	}
	defer ch.Close()

	//3. 利用channel声明要发送到的消息队列
	queue, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("declare a queue failed, err:%v\n", err)
		return
	}

	//4. 将消息发布到声明的队列
	body := bodyF(os.Args)
	err = ch.Publish(
		"",
		queue.Name,
		false, //立即
		false, //强制
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 持久
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		fmt.Printf("publish a message failed, err:%v\n", err)
		return
	}
	log.Printf("[x] Sent: %s", body)
}

func bodyF(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
