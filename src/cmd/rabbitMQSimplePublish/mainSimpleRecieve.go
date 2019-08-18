package rabbitMQSimplePublish

import "RabbitMQ"

func main(){
	rabbitmq := RabbitMQ.NewRabbitMQSimple("aigo")
	rabbitmq.ConsumeSimple()

}
