// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatGroupMsg is the golang structure of table gw_chat_group_msg for DAO operations like Where/Data.
type ChatGroupMsg struct {
	g.Meta      `orm:"table:gw_chat_group_msg, do:true"`
	Id          interface{} // 主键ID
	GroupId     interface{} // 群id
	SendUserId  interface{} // 发送者用户id
	MsgType     interface{} // 消息内容 1-文字 2-图片 3-视频
	Contents    interface{} // 消息内容
	Millisecond interface{} // 毫秒时间戳
	CreateTime  *gtime.Time // 创建时间
}
