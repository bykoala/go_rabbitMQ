package rabbitMQWork

import (
	"RabbitMQ"
	"strconv"
	"time"
)

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQSimple(""+"work_aigo")
	for i:=0;i<100;i++{
		rabbitmq.PublishSimple("hello,work_rabbitmq" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		println(i)
	}
}
