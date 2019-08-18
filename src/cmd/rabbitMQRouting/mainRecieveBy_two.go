package main

import "RabbitMQ"

func main(){
	by_two := RabbitMQ.NewRabbitMQRouting("ex_routing","by_two")
	by_two.RecieveRouting()
}
