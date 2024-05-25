package easychat

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	socketio "github.com/googollee/go-socket.io"
	"github.com/syyongx/php2go"
	"gweb/internal/dao"
	"time"
)

var GroupMsgRecordKeysFix = "group_msg_record_keys_fix_"

// 群消息
type GroupMessage struct {
	UserId  string `json:"user_id"`
	MsgType int    `json:"msg_type"`
	Content string `json:"content"`
}

// 发送消息给朋友
func (this *EasyChatServiceStruct) SendToGroup(s socketio.Conn, message string) string {
	//this.Log.Println("群消息：" + message)
	m := this.DecodeGroupMsg(message)
	//获取当前用户
	currUser, _ := this.GetUserBySocketId(s)
	msgMap := this.BuildGroupMessage(m, currUser.UserId)
	msg := this.Result.Success(g.Map{"message": msgMap})
	//通过遍历发送到群的每个成员
	this.GroupPush(m, currUser, msg, msgMap)
	//直接进库,后续再走MQ
	this.InGroupDb(msgMap)
	//进入到队列
	return msg
}

// 单聊数据直接入库
func (this *EasyChatServiceStruct) InGroupDb(recordMap g.Map) {
	_, err := dao.ChatGroupMsg.Ctx(gctx.New()).Insert(recordMap)
	if err != nil {
		g.Log().Warning(gctx.New(), err)
	}
}

// 构建消息体
func (this *EasyChatServiceStruct) BuildGroupMessage(m GroupMessage, userId string) g.Map {
	timestamp := time.Now().UnixNano() / 1e6
	if userId == "" {
		userId = m.UserId
	}
	dm := this.Dogs()
	userInfo := dm[userId]
	//获取用户信息
	return g.Map{
		"send_user_id": userId,
		"avatar":       userInfo.Avatar,
		"nickname":     userInfo.Nickname,
		"msg_type":     m.MsgType,
		"contents":     m.Content,
		"millisecond":  timestamp,
		"create_time":  gtime.Datetime(),
	}
}

// 发起极光推送
func (this *EasyChatServiceStruct) GroupPush(m GroupMessage, user Person, msg string, msgMap g.Map) {
	m.Content = this.GetMessageTypeText(m.MsgType, m.Content)
	mqScore := gconv.Int64(msgMap["millisecond"])
	targetUsersMap := this.Dogs()
	if len(targetUsersMap) > 0 {
		for _, v := range targetUsersMap {
			targetUid := gconv.String(v.Id)
			//判断用户是否在线,在线的不用发送
			u, err := this.GetSocketByUserId(targetUid)
			//无论如何都进入到离线队列中，对于收到ACK的消息才从离线队列中剔除
			mqKey := this.OfflineQueueFixMap.OffLineGroupMsgQueue + targetUid
			this.InQueuePre(mqKey, mqScore, msg)
			if err != nil {
			} else {
				//在线的直接发送消息
				this.SendMsg(u.Socket, this.Events.GroupReplyEvent, msg, mqKey, mqScore)
			}
		}
	}
}

// 解码群消息
func (this *EasyChatServiceStruct) DecodeGroupMsg(message string) GroupMessage {
	gm := GroupMessage{}
	err := php2go.JSONDecode([]byte(message), &gm)
	if err != nil {
		g.Dump(message)
		g.Log().Print(nil, err)
	}
	return gm
}

// 统一发送消息
func (this *EasyChatServiceStruct) SendMsg(s socketio.Conn, eventStr, msg, mqKey string, mqScore int64) {
	s.Emit(eventStr, msg, func(d string) {
		g.Dump("收到的ACK：" + d)
		this.EasyRedis.ZremRangByScore(mqKey, mqScore, mqScore)
	})
}

// ACK监听结果，如果没收到ACK则自动断开链接
func (this *EasyChatServiceStruct) ListenAck(s socketio.Conn, isClose chan bool) {
	//堵塞等到chan的结果
	res, ok := <-isClose
	if ok && res {
		this.Offline(s)
	}
	//defer close(isClose)
}

// 进入队列预处理
func (this *EasyChatServiceStruct) InQueuePre(mqKey string, mqScore int64, msg string) {
	this.EasyRedis.Zadd(mqKey, mqScore, msg)
	//获取队列长度，仅保留最新的100条记录
	queueLen := this.EasyRedis.Zcard(mqKey)
	if queueLen > 100 {
		delLen := queueLen - 100 - 1
		this.EasyRedis.ZremRangeByRank(mqKey, 0, delLen)
	}
}
