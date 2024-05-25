package users

import "github.com/gogf/gf/v2/frame/g"

type RolesUsedReq struct {
	g.Meta `path:"/chat/users/roles/:id" method:"put" tags:"UserService" summary:"设置角色占用情况"`
	IsUsed int `json:"is_used"`
}

type RolesUsedRes struct {
}
