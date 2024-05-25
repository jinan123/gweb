// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatUsers is the golang structure for table chat_users.
type ChatUsers struct {
	Id         int         `json:"id"         orm:"id"          description:"用户id"`
	Nickname   string      `json:"nickname"   orm:"nickname"    description:"昵称"`
	Avatar     string      `json:"avatar"     orm:"avatar"      description:"头像"`
	Sex        int         `json:"sex"        orm:"sex"         description:"性别 1-男 2-女 3-未知"`
	IsUsed     int         `json:"isUsed"     orm:"is_used"     description:"是否已被占用 1-是 0-否"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
}
