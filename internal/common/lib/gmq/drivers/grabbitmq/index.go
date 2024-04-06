package easyrabbitmq

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/streadway/amqp"
	"log"
)

//mq类
type RabbitMQ struct {
	//连接
	conn *amqp.Connection
	//管道
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key Simple模式 几乎用不到
	Key string
	//连接信息
	Mqurl string
	//日志
	Log *glog.Logger
}

//创建RabbitMQ结构体实例
func NewRabbitMQ(queuename string, exchange string, key string) *RabbitMQ {
	var err error
	configMap, err := g.Config("mq").Data(nil)
	if err != nil {
		panic(err)
	}
	url := gconv.String(configMap["Host"])
	rabbitmq := &RabbitMQ{
		QueueName: queuename,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     url,
		Log:       g.Log("mq"),
	}
	//创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		panic(err)
	}
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		panic(err)
	}
	return rabbitmq
}

//断开channel和connection
func (r *RabbitMQ) Destory() {
	if err := r.channel.Close(); err != nil {
		r.Log.Print(nil, err)
	}
	if err := r.conn.Close(); err != nil {
		r.Log.Print(nil, err)
	}
}

//处理错误函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
	}
}
