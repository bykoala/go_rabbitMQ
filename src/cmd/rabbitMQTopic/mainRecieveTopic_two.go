package main

import "RabbitMQ"

func main(){
	by_two := RabbitMQ.NewRabbitMQTopic("ex_topic","topic_two")
	by_two.RecieveTopic()
}
