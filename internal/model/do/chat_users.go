// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatUsers is the golang structure of table gw_chat_users for DAO operations like Where/Data.
type ChatUsers struct {
	g.Meta     `orm:"table:gw_chat_users, do:true"`
	Id         interface{} // 用户id
	Nickname   interface{} // 昵称
	Avatar     interface{} // 头像
	Sex        interface{} // 性别 1-男 2-女 3-未知
	IsUsed     interface{} // 是否已被占用 1-是 0-否
	CreateTime *gtime.Time // 创建时间
}
