package drivers

import "github.com/gogf/gf/v2/frame/g"

type RabbitMq struct {
}

func NewRabbitMq() *RabbitMq {
	return &RabbitMq{}
}

//设置配置文件
func (*RabbitMq) SetConfig(config g.Map) {

}

//入队
func (*RabbitMq) Enqueue(queueName, data string) {

}

//出队
func (*RabbitMq) Dequeue(queueName string) string {
	return ""
}

//队列长度
func (*RabbitMq) Len(queueName string) int {
	return 0
}
