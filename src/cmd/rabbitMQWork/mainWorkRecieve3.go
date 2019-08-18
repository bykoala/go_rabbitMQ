package rabbitMQWork

import "RabbitMQ"

func main(){
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "work_aigo")
	rabbitmq.ConsumeSimple()

}
