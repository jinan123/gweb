package easyredis

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

//向集合中添加一个或多个值
func (this *EasyRedisServiceStruct) Sadd(key string, value ...string) bool {
	val, _ := this.Handle().Do(nil, "SADD", gconv.Interfaces(append(g.SliceStr{key}, value...))...)
	return val.Bool()
}

//获取集合的数量
func (this *EasyRedisServiceStruct) Scard(key string) uint64 {
	val, _ := this.Handle().Do(nil, "SCARD", key)
	return val.Uint64()
}

//判断成员是否在集合中
func (this *EasyRedisServiceStruct) SisMember(key string, value interface{}) bool {
	val, _ := this.Handle().Do(nil, "SISMEMBER", key, value)
	return val.Bool()
}

//返回集合的所有成员
func (this *EasyRedisServiceStruct) SMembers(key string) []string {
	val, _ := this.Handle().Do(nil, "SMEMBERS", key)
	return val.Strings()
}

//移除并返回集合中的一个随机元素
func (this *EasyRedisServiceStruct) Spop(key string) string {
	val, _ := this.Handle().Do(nil, "SPOP", key)
	return val.String()
}

//移除集合中一个或多个成员
func (this *EasyRedisServiceStruct) Srem(key string, value ...string) bool {
	val, _ := this.Handle().Do(nil, "SREM", gconv.Interfaces(append(g.SliceStr{key}, value...))...)
	return val.Bool()
}

//返回第一个集合与其他集合之间的差异。求差集
func (this *EasyRedisServiceStruct) Sdiff(key string, keys ...string) []string {
	val, _ := this.Handle().Do(nil, "SDIFF", gconv.Interfaces(append(g.SliceStr{key}, keys...))...)
	return val.Strings()
}

//返回给定所有集合的交集
func (this *EasyRedisServiceStruct) Sinter(key string, keys ...string) []string {
	val, _ := this.Handle().Do(nil, "SINTER", gconv.Interfaces(append(g.SliceStr{key}, keys...))...)
	return val.Strings()
}
