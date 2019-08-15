package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url格式 amqp://账户;密码@RabbitMQ服务器地址:端口号/vhost
const MQURL = "amqp://by_name_root:byroot@127.0.0.1:5672/vir_root"

type RabbitMQ struct{
	conn *amqp.Connection
	channel *amqp.Channel
	//队列名词
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
}

//创建RabbitMQ结构体实例
func NewRabbitMQ(queueName,exchange,key string) *RabbitMQ{
	rabbitmq := &RabbitMQ{QueueName:queueName,Exchange:exchange,Key:key,Mqurl:MQURL}
	var err error
	//创建rabbitmq连接
	rabbitmq.conn,err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnErr(err,"创建连接错误！")
	rabbitmq.channel,err = rabbitmq.conn.Channel()
	rabbitmq.FailOnErr(err,"获取channel失败")
	return rabbitmq
}

//断开connection和channel
func (r *RabbitMQ) Destory(){
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) FailOnErr(err error,message string){
	if err != nil {
		log.Fatal("%s:%s",message,err)
		panic(fmt.Sprint("%s:%s",message,err))
	}
}

//创建简单模式Step:1.创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ{
	return NewRabbitMQ(queueName,"","")
}

//简单模式Step:2.简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string){
	//1.申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	//保证队列存在，消息能发送到队列中
	_,err := r.channel.QueueDeclare(r.QueueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil)
	if err != nil {
		fmt.Println(err)
	}
	//2.发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据exhange类型和routkey规则，如果无法找到符合条件的队列，那么会把发送的消息返回给消息发送者
		false,
		//如果为true，当exchan发送消息到队列后发现队列上没有绑定消费者，则会把消息返回给发送者
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(message),
		})
}