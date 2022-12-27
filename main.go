package main

import (
	mq "learn_code/rabbit_mq"
)

func main() {
	mq.Send("hello", "hello world")
	mq.Receive("hello")
}
