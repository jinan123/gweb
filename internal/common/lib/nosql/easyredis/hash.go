package easyredis

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"
)

//设置哈希表中的值
func (this *EasyRedisServiceStruct) Hset(key, fieldName, value interface{}) bool {
	result, _ := this.Handle().Do(nil, "HSET", key, fieldName, value)
	return result.Bool()
}

//获取哈希表中的值
func (this *EasyRedisServiceStruct) Hget(key, fieldName string) string {
	val, _ := this.Handle().Do(nil, "HGET", key, fieldName)
	return val.String()
}

//获取哈希表所有的中
func (this *EasyRedisServiceStruct) HGetAll(key string) g.Map {
	val, _ := this.Handle().Do(nil, "HGETALL", key)
	return this.StrArrayToMap(val.Vars())
}

//获取哈希表所有的字段
func (this *EasyRedisServiceStruct) Hkeys(key string) []string {
	val, _ := this.Handle().Do(nil, "HKEYS", key)
	return val.Strings()
}

//获取更多字段
func (this *EasyRedisServiceStruct) Hmget(key interface{}, field ...interface{}) []string {
	val, _ := this.Handle().Do(nil, "HMGET", append(g.Slice{key}, field...)...)
	return val.Strings()
}

//同时设置多个值
func (this *EasyRedisServiceStruct) Hmset(key string, value ...g.Map) []string {
	s := g.Slice{key}
	for _, v := range value {
		s = append(g.Slice{key}, gutil.MapToSlice(v)...)
	}
	val, _ := this.Handle().Do(nil, "HMSET", s...)
	return val.Strings()
}

//删除哈希表中的值
func (this *EasyRedisServiceStruct) Hdel(key, fieldName string) bool {
	result, _ := this.Handle().Do(nil, "HDEL", key, fieldName)
	return result.Bool()
}
