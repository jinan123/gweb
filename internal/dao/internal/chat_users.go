// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatUsersDao is the data access object for table gw_chat_users.
type ChatUsersDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns ChatUsersColumns // columns contains all the column names of Table for convenient usage.
}

// ChatUsersColumns defines and stores column names for table gw_chat_users.
type ChatUsersColumns struct {
	Id         string // 用户id
	Nickname   string // 昵称
	Avatar     string // 头像
	Sex        string // 性别 1-男 2-女 3-未知
	IsUsed     string // 是否已被占用 1-是 0-否
	CreateTime string // 创建时间
}

// chatUsersColumns holds the columns for table gw_chat_users.
var chatUsersColumns = ChatUsersColumns{
	Id:         "id",
	Nickname:   "nickname",
	Avatar:     "avatar",
	Sex:        "sex",
	IsUsed:     "is_used",
	CreateTime: "create_time",
}

// NewChatUsersDao creates and returns a new DAO object for table data access.
func NewChatUsersDao() *ChatUsersDao {
	return &ChatUsersDao{
		group:   "default",
		table:   "gw_chat_users",
		columns: chatUsersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChatUsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChatUsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChatUsersDao) Columns() ChatUsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChatUsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChatUsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChatUsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
