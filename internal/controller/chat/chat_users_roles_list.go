package chat

import (
	"context"
	"gweb/internal/dao"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gweb/api/chat/users"
)

// RolesList 获取狗狗列表
func (c *ControllerUsers) RolesList(ctx context.Context, req *users.RolesListReq) (res *users.RolesListRes, err error) {
	result, err := dao.ChatUsers.Ctx(ctx).Where("is_used = 0").Order("id asc").All()
	if err != nil {
		return nil, err
	}
	dogs := []users.Dogs{}
	if result.IsEmpty() {
		return &users.RolesListRes{Dogs: dogs}, nil
	}
	err = result.Structs(&dogs)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeNotImplemented)
	}
	return &users.RolesListRes{Dogs: dogs}, nil
}
