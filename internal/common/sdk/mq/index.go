package mq

import "github.com/gogf/gf/v2/frame/g"

//定义mq接口
type EasyMq interface {
	SetConfig(config g.Map)
	//入队
	Enqueue(queueName, data string)
	//出队
	Dequeue(queueName string) string
	//队列长度
	Len(queueName string) int
}
