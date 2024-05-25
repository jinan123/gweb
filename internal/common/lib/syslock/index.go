package syslock

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/syyongx/php2go"
	"gweb/internal/common/lib/nosql/easyredis"
	"time"
)

var EasyLockService = EasyLockServiceStruct{}

type EasyLockServiceStruct struct {
	Key         string
	Value       string
	TimeOut     int
	RedisHandle *easyredis.EasyRedisServiceStruct
	IsGetLock   bool
}

// 初始化锁
func New(key string) *EasyLockServiceStruct {
	return &EasyLockServiceStruct{
		Key:         key,
		Value:       php2go.Uniqid("sys_lock_") + gconv.String(php2go.Rand(100, 999)),
		TimeOut:     10,
		RedisHandle: easyredis.New(),
	}
}

// 带参数初始化
func NewWithParams(key, value string, timeOut int) *EasyLockServiceStruct {
	return &EasyLockServiceStruct{
		Key:         key,
		Value:       value,
		TimeOut:     timeOut,
		RedisHandle: easyredis.New(),
	}
}

// 上锁
func (this *EasyLockServiceStruct) Lock() bool {
	return this.RedisHandle.SetNx(this.Key, this.Value, this.TimeOut)
}

// 解锁
func (this *EasyLockServiceStruct) UnLock() bool {
	const unlockScript = "if redis.call('get',KEYS[1]) == ARGV[1] then return redis.call('del',KEYS[1]) else return 0 end"
	result, _ := this.RedisHandle.Handle().Do(nil, "EVAL", unlockScript, 1, this.Key, this.Value)
	return result.Bool()
}

// 自动抢夺锁头
func (this *EasyLockServiceStruct) AutoLock() *EasyLockServiceStruct {
	start := gtime.Timestamp()
	for {
		//抢夺锁头
		if this.Lock() == true {
			this.IsGetLock = true
			return this
		}
		//每10毫秒获取一次锁头
		time.Sleep(50 * time.Millisecond)
		//如果当前时间已经超过最长抢夺时间，则释放锁头
		if gtime.Timestamp()-start >= int64(this.TimeOut) {
			this.UnLock()
			this.IsGetLock = false
			return this
		}
	}
}
