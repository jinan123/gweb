package users

import "github.com/gogf/gf/v2/frame/g"

//获取聊天记录
type MsgRecordReq struct {
	g.Meta      `path:"/chat/record" method:"get" tags:"UserService" summary:"获取聊天记录"`
	Millisecond int64 `json:"millisecond"`
}

//聊天记录模型
type MsgRecords struct {
	SendUserId  int    `json:"send_user_id"`
	MsgType     int    `json:"msg_type"`
	Contents    string `json:"contents"`
	Millisecond int64  `json:"millisecond"`
	Create_time string `json:"create_time"`
}

//获取聊天记录响应
type MsgRecordRes struct {
	Records        []MsgRecords `json:"records"`
	MinMillisecond int64        `json:"min_millisecond"`
}

//聊天记录模型
type SendGiftMsgReq struct {
	g.Meta `path:"/chat/gift" method:"post" tags:"UserService" summary:"礼物"`
	Uid    string `json:"uid"`
	Gtype  int    `json:"gtype"`
}

//获取聊天记录响应
type SendGiftMsgRes struct {
}
