package easychat

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	socketio "github.com/googollee/go-socket.io"
	"github.com/syyongx/php2go"
	"gweb/api/chat/users"
	"gweb/internal/common/lib/syslock"
	"gweb/internal/dao"
	"time"
)

type SocketCtx struct {
	LastTime int64
	UserId   string
}

// 获取Ctx内容
func (this *EasyChatServiceStruct) GetSocketCtxInfo(s socketio.Conn) SocketCtx {
	ctx := s.Context()
	if ctx == nil {
		return SocketCtx{}
	}
	return ctx.(SocketCtx)
}

// 设置Ctx内容
func (this *EasyChatServiceStruct) SetSocketCtxInfo(s socketio.Conn, content SocketCtx) {
	s.SetContext(content)
}

// 上线
func (this *EasyChatServiceStruct) Online(s socketio.Conn, msg string) string {
	msgMap := g.Map{}
	err := php2go.JSONDecode([]byte(msg), &msgMap)
	if err != nil {
		return this.Result.Unauthorized("缺失token")
	}
	uid := gconv.String(msgMap["id"])
	this.SetSocketCtxInfo(s, SocketCtx{UserId: uid, LastTime: 0})
	WriteUserSocketMap(gconv.String(uid), Person{Socket: s})
	this.OnLineDogs(uid)
	go func() {
		all := ReadAllSocketMap()
		for _, v := range all {
			v.Emit("reloadDogs", "doReload", func(data string) {
				g.Dump("-----收到的ack----")
				g.Dump(data)
			})
		}
	}()
	return this.Result.SuccessMsg()
}

// 下线
func (this *EasyChatServiceStruct) Offline(s socketio.Conn) string {
	if s != nil {
		ctxInfo := this.GetSocketCtxInfo(s)
		user, err := ReadUserSocketMap(ctxInfo.UserId)
		if err != nil {

		} else {
			if user.Socket != nil {
				//只有当前的socketID相等的时候才处理下线
				if user.Socket.ID() == s.ID() {
					RemoveUserSocketMap(ctxInfo.UserId)
					s.SetContext(SocketCtx{UserId: "", LastTime: php2go.Time()})
				} else {
					//g.Log().Print(gctx.New(), user.Socket.ID()+"-----------离线ID不等---------"+socketId)
				}
			}

		}
	}
	DelAllSocketMap(s)
	return this.Result.SuccessMsg()
}

// 拉取聊天未读消息
func (this *EasyChatServiceStruct) UnRead(s socketio.Conn) string {
	//只有抢夺到锁的携程才可以发送离线消息,避免ACK未执行重复推送离线消息
	lockKey := "socketio_unread_" + s.ID()
	cLock := syslock.New(lockKey).AutoLock()
	if cLock.IsGetLock {
		defer cLock.UnLock()
		//获得锁 计算上次的操作时间
		sCtx := this.GetSocketCtxInfo(s)
		if gtime.Timestamp()-sCtx.LastTime > 1 {
			ctxInfo := this.GetSocketCtxInfo(s)
			if ctxInfo.UserId != "" {
				userId := ctxInfo.UserId
				time.Sleep(100 * time.Millisecond)
				this.UnReadGroup(s, userId)
			}
			sCtx.LastTime = gtime.Timestamp()
			this.SetSocketCtxInfo(s, sCtx)
		}
	}
	return this.Result.SuccessMsg()
}

// 群未读
func (this *EasyChatServiceStruct) UnReadGroup(s socketio.Conn, userId string) {
	//推送好友未读消息
	currTime := gtime.Now().Timestamp() * 1000
	offLineGroupMsgQueueKey := this.OfflineQueueFixMap.OffLineGroupMsgQueue + userId
	//通过延时队列实现延时发送
	msgs := this.EasyRedis.Zrangebyscore(offLineGroupMsgQueueKey, 0, currTime)
	//拉取离线队列
	if len(msgs) > 0 {
		for _, msg := range msgs {
			msgMap := g.Map{}
			err := php2go.JSONDecode([]byte(msg), &msgMap)
			if err != nil {
				g.Log().Warning(nil, err)
			} else {
				data := gconv.Map(msgMap["data"])
				msgDataMap := gconv.Map(data["message"])
				mqScore := gconv.Int64(msgDataMap["millisecond"])
				this.SendMsg(s, this.Events.GroupReplyEvent, msg, offLineGroupMsgQueueKey, mqScore)
			}
		}
	}
}

//获取狗子列表
func (this *EasyChatServiceStruct) Dogs() map[string]users.Dogs {
	result, err := dao.ChatUsers.Ctx(gctx.New()).Cache(gdb.CacheOption{
		Duration: 10 * time.Hour,
		Name:     "dogs",
	}).Order("id asc").All()
	dm := map[string]users.Dogs{}
	dogs := []users.Dogs{}
	if err != nil {
		g.Log().Warning(gctx.New(), err)
		return dm
	}

	if result.IsEmpty() {
		return dm
	}
	err = result.Structs(&dogs)
	if err != nil {
		g.Log().Warning(gctx.New(), err)
		return dm
	}
	for _, v := range dogs {
		dm[gconv.String(v.Id)] = v
	}
	return dm
}

//获取狗子列表
func (this *EasyChatServiceStruct) OnLineDogs(id string) {
	_, err := dao.ChatUsers.Ctx(gctx.New()).Where("id = ?", id).Update(g.Map{"is_used": 1})
	if err != nil {
		g.Log().Warning(gctx.New(), err)
	}
}
