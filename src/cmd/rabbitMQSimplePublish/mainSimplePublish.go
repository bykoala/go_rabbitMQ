package rabbitMQSimplePublish

import (
	"RabbitMQ"
	"fmt"
)

func main(){
	rabbitmq := RabbitMQ.NewRabbitMQSimple("aigo")
	rabbitmq.PublishSimple("say hello")
	fmt.Println("发送成功")
}
