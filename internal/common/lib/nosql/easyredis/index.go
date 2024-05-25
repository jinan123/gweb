package easyredis

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gomodule/redigo/redis"
)

var EasyRedisService = EasyRedisServiceStruct{}

type SubscribeCallback func(channel, message string)

type EasyRedisServiceStruct struct {
	//handle               *gredis.Redis
	SubscribeCallBackMap map[string]SubscribeCallback
	PubSubConn           redis.PubSubConn
	//Pool                 gredis.RedisConn
}

// 初始化redis
func New() *EasyRedisServiceStruct {
	//注意，此处会发生连接超时，也可以设置超时为永不超时，但是严谨起见应该直接用redis连接池
	//handle := g.Redis()
	return &EasyRedisServiceStruct{
		//handle:               handle,
		SubscribeCallBackMap: make(map[string]SubscribeCallback),
		//PubSubConn:           redis.PubSubConn{Conn: handle.Conn()},
	}
}

// 获取redis操作句柄
func (this *EasyRedisServiceStruct) Handle() *gredis.Redis {
	return g.Redis()
}

// 关闭链接
func (this *EasyRedisServiceStruct) Close() error {
	return this.Handle().Close(nil)
}

// GET
func (this *EasyRedisServiceStruct) Get(key string) string {
	v, _ := this.Handle().Do(nil, "GET", key)
	return v.String()
}

// Set
func (this *EasyRedisServiceStruct) Set(key, value string, timeOut ...int) {
	this.Handle().Do(nil, "SET", key, value)
	if len(timeOut) > 0 && timeOut[0] > 0 {
		this.Expire(key, timeOut[0])
	}

}

// SetNx
func (this *EasyRedisServiceStruct) SetNx(key, value string, timeOut int) bool {
	result, err := this.Handle().Do(nil, "SET", key, value, "NX", "EX", timeOut)
	if err != nil {
		return false
	}
	return result.Bool()
}

// DEl
func (this *EasyRedisServiceStruct) Del(key string) {
	this.Handle().Do(nil, "DEL", key)
}

// 判断key是否存在 1-存在 0-不存在
func (this *EasyRedisServiceStruct) Exists(key string) int {
	has, _ := this.Handle().Do(nil, "EXISTS", key)
	return gconv.Int(has)
}

// 设置有效期
func (this *EasyRedisServiceStruct) Expire(key string, timeOut int) bool {
	result, _ := this.Handle().Do(nil, "EXPIRE", key, timeOut)
	return result.Bool()
}

// 返回值的有效期
func (this *EasyRedisServiceStruct) Ttl(key string) int {
	val, _ := this.Handle().Do(nil, "TTL", key)
	return val.Int()
}

// 返回keys
func (this *EasyRedisServiceStruct) Keys(key ...string) []string {
	keyword := "*"
	if len(key) > 0 {
		keyword = key[0]
	}
	val, _ := this.Handle().Do(nil, "keys", keyword)
	return val.Strings()
}

// 字符串数组转map
func (this *EasyRedisServiceStruct) StrArrayToMap(arr []*g.Var) g.Map {
	arrLen := len(arr)
	result := g.Map{}
	if arrLen == 0 {
		return result
	}
	for i := 0; i < arrLen; {
		result[arr[i].String()] = arr[i+1]
		i = i + 2
	}
	return result
}
