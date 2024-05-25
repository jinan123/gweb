package easychat

//定义队列字典
type QueueMap struct {
	GroupMsgQueue string //群聊记录消息队列
}

//推送事件定义
type EventsMap struct {
	ReplyEvent      string //回复
	SendEvent       string //发送
	LoginEvent      string //登录
	OutEvent        string //退出
	GroupReplyEvent string //群消息推送
}

//离线队列
type OffLineQueueFixMap struct {
	OffLineGroupMsgQueue string //群聊离线消息队列 (离线重发)
}

//获取队列字典
func GetQueueMap() QueueMap {
	return QueueMap{
		GroupMsgQueue: "groupMsgQueue",
	}
}

//获取事件字典
func GetEventsMap() EventsMap {
	return EventsMap{
		ReplyEvent:      "replyEvent",
		SendEvent:       "sendEvent",
		LoginEvent:      "loginEvent",
		OutEvent:        "outEvent",
		GroupReplyEvent: "groupReplyEvent",
	}
}

//离线队列
func GetOfflineQueueFixMap() OffLineQueueFixMap {
	return OffLineQueueFixMap{
		OffLineGroupMsgQueue: "offLineGroupZMsgQueue_",
	}
}
