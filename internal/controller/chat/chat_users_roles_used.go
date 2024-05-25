package chat

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gweb/api/chat/users"
)

func (c *ControllerUsers) RolesUsed(ctx context.Context, req *users.RolesUsedReq) (res *users.RolesUsedRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
