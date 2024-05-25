// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatGroupMsg is the golang structure for table chat_group_msg.
type ChatGroupMsg struct {
	Id          int         `json:"id"          orm:"id"           description:"主键ID"`
	GroupId     int         `json:"groupId"     orm:"group_id"     description:"群id"`
	SendUserId  int         `json:"sendUserId"  orm:"send_user_id" description:"发送者用户id"`
	MsgType     int         `json:"msgType"     orm:"msg_type"     description:"消息内容 1-文字 2-图片 3-视频"`
	Contents    string      `json:"contents"    orm:"contents"     description:"消息内容"`
	Millisecond int64       `json:"millisecond" orm:"millisecond"  description:"毫秒时间戳"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:"创建时间"`
}
