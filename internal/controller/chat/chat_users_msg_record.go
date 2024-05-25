package chat

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gweb/api/chat/users"
)

func (c *ControllerUsers) MsgRecord(ctx context.Context, req *users.MsgRecordReq) (res *users.MsgRecordRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
