package users

import "github.com/gogf/gf/v2/frame/g"

type RolesListReq struct {
	g.Meta `path:"/chat/users/roles" method:"get" tags:"UserService" summary:"可用角色列表"`
}
type Dogs struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
type RolesListRes struct {
	Dogs []Dogs `json:"dogs"`
}
