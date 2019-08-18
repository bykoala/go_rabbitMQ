package main

import (
	"RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main()  {
	by_one := RabbitMQ.NewRabbitMQTopic("ex_topic","topic_one")
	by_two := RabbitMQ.NewRabbitMQTopic("ex_topic","topic_two")
	for i:= 0;i <= 10;i++{
		by_one.PublishTopic("hello topic one" + strconv.Itoa(i))
		by_two.PublishTopic("hello topic two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
