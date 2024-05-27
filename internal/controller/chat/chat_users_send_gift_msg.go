package chat

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"gweb/internal/common/lib/nosql/easyredis"

	"gweb/api/chat/users"
)

func (c *ControllerUsers) SendGiftMsg(ctx context.Context, req *users.SendGiftMsgReq) (res *users.SendGiftMsgRes, err error) {
	rd := easyredis.New()
	rd.Lpush("giftMsgQueue", gconv.String(g.Map{"uid": req.Uid, "gtype": req.Gtype}))
	return
}
