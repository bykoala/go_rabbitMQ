package main

import "RabbitMQ"

func main(){
	by_one := RabbitMQ.NewRabbitMQRouting("ex_routing","by_one")
	by_one.RecieveRouting()
}
