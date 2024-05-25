package easychat

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	socketio "github.com/googollee/go-socket.io"
	"gweb/internal/common/lib/nosql/easyredis"
	"gweb/internal/common/sdk/easychat/result"
	"gweb/internal/common/sdk/mq"
	"gweb/internal/common/sdk/mq/drivers"
)

var EasychatService = EasyChatServiceStruct{}

//聊天模型
type EasyChatServiceStruct struct {
	Result             result.ResultServiceStruct
	Events             EventsMap
	Mq                 mq.EasyMq
	QueueNames         QueueMap
	OfflineQueueFixMap OffLineQueueFixMap
	Log                *glog.Logger
	EasyRedis          *easyredis.EasyRedisServiceStruct
}

//启动聊天
func New() *EasyChatServiceStruct {
	//注册服务
	return &EasyChatServiceStruct{
		Result:             result.ResultService,
		Events:             GetEventsMap(),
		Mq:                 drivers.NewRedisMq(),
		QueueNames:         GetQueueMap(),
		Log:                g.Log(),
		OfflineQueueFixMap: GetOfflineQueueFixMap(),
		EasyRedis:          easyredis.New(),
	}
}

//获取某个在线用户
func (this *EasyChatServiceStruct) GetUserBySocketId(s socketio.Conn) (Person, error) {
	ctxInfo := this.GetSocketCtxInfo(s)
	return this.GetSocketByUserId(ctxInfo.UserId)
}

//获取某个用户的socket
func (this *EasyChatServiceStruct) GetSocketByUserId(uid string) (Person, error) {
	user, err := ReadUserSocketMap(uid)
	if err != nil {
		return Person{}, gerror.New("用户未在线")
	}
	return user, nil
}

//获取在线人数
func (this *EasyChatServiceStruct) OnlineCount() int {
	i := 0
	SocketUserIdMap.Range(func(key, value interface{}) bool {
		i++
		return true
	})
	return i
}

//返回对应的语音信息
func (this *EasyChatServiceStruct) GetMessageTypeText(ty int, c string) string {
	mt := map[int]string{
		1: c,
		2: "【图片】",
		3: "【视频】",
		4: "【语音】",
	}
	return mt[ty]
}
