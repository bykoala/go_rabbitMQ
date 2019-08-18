package main

import "RabbitMQ"

func main(){
	by_one := RabbitMQ.NewRabbitMQTopic("ex_topic","#")
	by_one.RecieveTopic()
}
