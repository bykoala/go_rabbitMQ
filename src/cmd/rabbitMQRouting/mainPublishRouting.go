package main

import (
	"RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main()  {
	by_one := RabbitMQ.NewRabbitMQRouting("ex_routing","by_one")
	by_two := RabbitMQ.NewRabbitMQRouting("ex_routing","by_two")
	for i:= 0;i <= 10;i++{
		by_one.PublishRouting("hello routing one" + strconv.Itoa(i))
		by_two.PublishRouting("hello routing two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
