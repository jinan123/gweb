package easyredis

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gomodule/redigo/redis"
	"time"
	"unsafe"
)

func (this *EasyRedisServiceStruct) GetPubSubConn() *EasyRedisServiceStruct {
	//this.Pool , err := this.Handle().Conn(nil)
	//this.PubSubConn = redis.PubSubConn{Conn: this.Pool.Receive(nil)}
	return this
}

// 注册订阅者
func (this *EasyRedisServiceStruct) RegisterSubScribeHandle(channel interface{}, cb SubscribeCallback) {
	err := this.PubSubConn.Subscribe(channel)
	if err != nil {
		g.Log().Print(nil, err)
	}
	this.SubscribeCallBackMap[gconv.String(channel)] = cb
}

// 订阅
func (this *EasyRedisServiceStruct) StartSubScribe() {
	go func() {
		for {
			switch res := this.PubSubConn.Receive().(type) {
			case redis.Message:
				channel := (*string)(unsafe.Pointer(&res.Channel))
				message := (*string)(unsafe.Pointer(&res.Data))
				this.SubscribeCallBackMap[*channel](*channel, *message)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", res.Channel, res.Kind, res.Count)
			case error:
				g.Log().Print(nil, "SubScribe error")
				continue
			}
		}
		//this.Pool.Close(nil)
	}()
	//阻止线程结束
	for {
		time.Sleep(1 * time.Second)
	}
}

// 发布
func (this *EasyRedisServiceStruct) PubLish(key, data string) {
	this.Handle().Do(nil, "PUBLISH", key, data)
}
