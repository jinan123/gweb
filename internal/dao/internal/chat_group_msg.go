// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatGroupMsgDao is the data access object for table gw_chat_group_msg.
type ChatGroupMsgDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ChatGroupMsgColumns // columns contains all the column names of Table for convenient usage.
}

// ChatGroupMsgColumns defines and stores column names for table gw_chat_group_msg.
type ChatGroupMsgColumns struct {
	Id          string // 主键ID
	GroupId     string // 群id
	SendUserId  string // 发送者用户id
	MsgType     string // 消息内容 1-文字 2-图片 3-视频
	Contents    string // 消息内容
	Millisecond string // 毫秒时间戳
	CreateTime  string // 创建时间
}

// chatGroupMsgColumns holds the columns for table gw_chat_group_msg.
var chatGroupMsgColumns = ChatGroupMsgColumns{
	Id:          "id",
	GroupId:     "group_id",
	SendUserId:  "send_user_id",
	MsgType:     "msg_type",
	Contents:    "contents",
	Millisecond: "millisecond",
	CreateTime:  "create_time",
}

// NewChatGroupMsgDao creates and returns a new DAO object for table data access.
func NewChatGroupMsgDao() *ChatGroupMsgDao {
	return &ChatGroupMsgDao{
		group:   "default",
		table:   "gw_chat_group_msg",
		columns: chatGroupMsgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChatGroupMsgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChatGroupMsgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChatGroupMsgDao) Columns() ChatGroupMsgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChatGroupMsgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChatGroupMsgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChatGroupMsgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
