package easyredis

//将一个或多个值插入到列表头部
func (this *EasyRedisServiceStruct) Lpush(key, value string) {
	this.Handle().Do(nil, "LPUSH", key, value)
}

//将一个或多个值插入到列表尾部
func (this *EasyRedisServiceStruct) Rpush(key, value string) {
	this.Handle().Do(nil, "RPUSH", key, value)
}

//移出并获取列表的第一个元素
func (this *EasyRedisServiceStruct) Lpop(key string) string {
	val, _ := this.Handle().Do(nil, "LPOP", key)
	return val.String()
}

//移除列表的最后一个元素
func (this *EasyRedisServiceStruct) Rpop(key string) string {
	val, _ := this.Handle().Do(nil, "RPOP", key)
	return val.String()
}

//返回队列长度
func (this *EasyRedisServiceStruct) Llen(key string) int {
	val, _ := this.Handle().Do(nil, "LLEN", key)
	return val.Int()
}

//通过索引获取列表中的元素
func (this *EasyRedisServiceStruct) Lindex(key string, index int) string {
	val, _ := this.Handle().Do(nil, "LINDEX", key, index)
	return val.String()
}
