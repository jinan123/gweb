// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chat

import (
	"context"

	"gweb/api/chat/users"
)

type IChatUsers interface {
	MsgRecord(ctx context.Context, req *users.MsgRecordReq) (res *users.MsgRecordRes, err error)
	SendGiftMsg(ctx context.Context, req *users.SendGiftMsgReq) (res *users.SendGiftMsgRes, err error)
	RolesList(ctx context.Context, req *users.RolesListReq) (res *users.RolesListRes, err error)
	RolesUsed(ctx context.Context, req *users.RolesUsedReq) (res *users.RolesUsedRes, err error)
}
