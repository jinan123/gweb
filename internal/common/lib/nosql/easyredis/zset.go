package easyredis

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// 向有序集中总添加值
func (this *EasyRedisServiceStruct) Zadd(key string, score interface{}, value string) bool {
	val, _ := this.Handle().Do(nil, "ZADD", key, score, value)
	return val.Bool()
}

// 从有序集合中获取值，正向排序
func (this *EasyRedisServiceStruct) Zrange(key string, start, end interface{}) []string {
	val, _ := this.Handle().Do(nil, "ZRANGE", key, start, end)
	return val.Strings()
}

// 从有序集合中获取值，反向排序
func (this *EasyRedisServiceStruct) Zrevrange(key string, start, end interface{}) []string {
	val, _ := this.Handle().Do(nil, "ZREVRANGE", key, start, end)
	return val.Strings()
}

// 获取指定value的score
func (this *EasyRedisServiceStruct) Zscore(key string, value string) float64 {
	val, _ := this.Handle().Do(nil, "ZSCORE", key, value)
	return val.Float64()
}

// 获取指定key的集合长度
func (this *EasyRedisServiceStruct) Zcard(key string) int {
	val, _ := this.Handle().Do(nil, "ZCARD", key)
	return val.Int()
}

// 根据分值区间获取值(分数从低到高)
func (this *EasyRedisServiceStruct) Zrangebyscore(key string, min, max interface{}) []string {
	val, _ := this.Handle().Do(nil, "ZRANGEBYSCORE", key, min, max)
	return val.Strings()
}

// 根据分值区间获取值(分数从高到低)
func (this *EasyRedisServiceStruct) ZrevrangeByScore(key string, max, min interface{}) []string {
	val, _ := this.Handle().Do(nil, "ZREVRANGEBYSCORE", key, max, min)
	return val.Strings()
}

// 分页返回
func (this *EasyRedisServiceStruct) ZrevrangeByScoreWithPage(key string, max, min, offset, count interface{}) []string {
	val, _ := this.Handle().Do(nil, "ZREVRANGEBYSCORE", key, max, min, "WITHSCORES", "limit", offset, count)
	return val.Strings()
}

// 根据分值区间获取值(带分数)
func (this *EasyRedisServiceStruct) ZrangebyscoreWithScore(key string, min, max interface{}) g.Map {
	val, _ := this.Handle().Do(nil, "ZRANGEBYSCORE", key, min, max, "WITHSCORES")
	return this.StrArrayToMap(val.Vars())
}

// 从有序集中删除的一个或多个值
func (this *EasyRedisServiceStruct) Zrem(key string, value ...string) bool {
	val, _ := this.Handle().Do(nil, "ZREM", gconv.Interfaces(append(g.SliceStr{key}, value...))...)
	return val.Bool()
}

// 从有序集中删除的一个或多个值(根据分数)
func (this *EasyRedisServiceStruct) ZremRangByScore(key string, min, max interface{}) bool {
	val, _ := this.Handle().Do(nil, "ZREMRANGEBYSCORE", key, min, max)
	return val.Bool()
}

// 删除分数区间的值（分数从小到大排序）
func (this *EasyRedisServiceStruct) ZremRangeByRank(key string, min, len int) bool {
	val, _ := this.Handle().Do(nil, "ZREMRANGEBYRANK", key, min, len)
	return val.Bool()
}
