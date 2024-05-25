package drivers

import (
	"github.com/gogf/gf/v2/frame/g"
	"gweb/internal/common/lib/nosql/easyredis"
)

type RedisEasyMq struct {
	handle *easyredis.EasyRedisServiceStruct
}

//返回redis驱动
func NewRedisMq() *RedisEasyMq {
	return &RedisEasyMq{
		handle: easyredis.New(),
	}
}

//设置配置文件
func (this *RedisEasyMq) SetConfig(config g.Map) {

}

//入队
func (this *RedisEasyMq) Enqueue(queueName, data string) {
	this.handle.Lpush(queueName, data)
}

//出队
func (this *RedisEasyMq) Dequeue(queueName string) string {
	return this.handle.Rpop(queueName)
}

//队列长度
func (this *RedisEasyMq) Len(queueName string) int {
	return this.handle.Llen(queueName)
}
