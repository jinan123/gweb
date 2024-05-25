package easychat

import (
	"github.com/gogf/gf/v2/errors/gerror"
	socketio "github.com/googollee/go-socket.io"
	"sync"
)

//个人聊天模型
type Person struct {
	Socket socketio.Conn
	UserId string
}

// 定义锁
var RLock sync.RWMutex

// 定义sockeit与用户id关系map
var SocketUserIdMap sync.Map

var UserSocketMapRelation map[string]Person = map[string]Person{}

//保存当前所有的socket链接
var AllSocketMap = map[string]socketio.Conn{}

// 写socket
func WriteAllSocketMap(s socketio.Conn) {
	RLock.RLock()
	defer RLock.RUnlock()
	AllSocketMap[s.ID()] = s
}

// 删socket
func DelAllSocketMap(s socketio.Conn) {
	RLock.RLock()
	defer RLock.RUnlock()
	delete(AllSocketMap, s.ID())
}

//读取 所有的socket
func ReadAllSocketMap() map[string]socketio.Conn {
	RLock.RLock()
	defer RLock.RUnlock()
	return AllSocketMap
}

// 通过用户id获取socket句柄
func ReadUserSocketMap(userId string) (Person, error) {
	RLock.RLock()
	defer RLock.RUnlock()
	p, ok := UserSocketMapRelation[userId]
	if ok {
		return p, nil
	}
	return Person{}, gerror.New("不存在的用户")
}

// 写入用户id与与socket句柄
func WriteUserSocketMap(userId string, p Person) {
	RLock.Lock()
	defer RLock.Unlock()
	UserSocketMapRelation[userId] = p
}

// 移除
func RemoveUserSocketMap(userId string) {
	RLock.Lock()
	defer RLock.Unlock()
	delete(UserSocketMapRelation, userId)
}

// 返回整个map
func OnlineMap() map[string]Person {
	RLock.RLock()
	defer RLock.RUnlock()
	s := UserSocketMapRelation
	return s
}
