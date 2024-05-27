package easychat

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/syyongx/php2go"
	"time"
)

//监听聊天NQ
func (this *EasyChatServiceStruct) ListenMq() {
	go this.MqGift()
}

//监听消息队列
func (this *EasyChatServiceStruct) MqGift() GroupMessage {
	this.Log.Print(nil, "start receive system data")
	gitfUrl := map[int]string{
		1: "https://tingting-1318709742.cos.ap-beijing.myqcloud.com/5f/1bb6f61a8133b2e92999152c08d756.gif?attname=66%E3%80%90%E9%B1%BC%E6%82%A6%E7%B4%A0%E6%9D%90%E3%80%91%E7%81%AB%E7%AE%AD6.gif",
		2: "https://tingting-1318709742.cos.ap-beijing.myqcloud.com/9c/bb06b7b1d40af3e6ad2d588a70c0a6.gif?attname=43%E3%80%90%E9%B1%BC%E6%82%A6%E7%B4%A0%E6%9D%90%E3%80%91%E8%93%9D%E7%8E%AB%E7%91%B0.gif",
		3: "https://tingting-1318709742.cos.ap-beijing.myqcloud.com/9b/3a67671235848b7ab1dc04e2f2de76.gif?attname=62%E3%80%90%E9%B1%BC%E6%82%A6%E7%B4%A0%E6%9D%90%E3%80%91%E7%88%B1%E5%BF%83.gif",
		4: "https://tingting-1318709742.cos.ap-beijing.myqcloud.com/62/dace20cae409a063cb0ce8662a1f2c.gif?attname=17%E3%80%90%E9%B1%BC%E6%82%A6%E7%B4%A0%E6%9D%90%E3%80%91%E7%94%9C%E7%94%9C%E5%9C%88.gif",
	}
	for {
		msg := this.Mq.Dequeue(this.QueueNames.Gift)
		if msg == "" {
			time.Sleep(time.Duration(1) * time.Second)
		} else {
			//解码数据
			apply := g.Map{}
			err := php2go.JSONDecode([]byte(msg), &apply)
			if err != nil {
				this.Log.Print(nil, err)
			} else {
				//拉取用户
				targetUsersMap := this.Dogs()
				tuid := gconv.Int(apply["uid"])
				if tuid > 0 {
					targetUid := gconv.String(tuid)
					//判断用户是否在线,在线的不用发送
					u, err := this.GetSocketByUserId(targetUid)
					if err != nil {
					} else {
						//在线的直接发送消息
						u.Socket.Emit("gift", gconv.String(g.Map{
							"url": gitfUrl[gconv.Int(apply["gtype"])],
						}), func(d string) {
							g.Dump("---礼物ack---")
						})
					}
				} else {
					if len(targetUsersMap) > 0 {
						for _, v := range targetUsersMap {
							targetUid := gconv.String(v.Id)
							//判断用户是否在线,在线的不用发送
							u, err := this.GetSocketByUserId(targetUid)
							if err != nil {
							} else {
								//在线的直接发送消息
								u.Socket.Emit("gift", gconv.String(g.Map{
									"url": gitfUrl[gconv.Int(apply["gtype"])],
								}), func(d string) {
									g.Dump("---礼物ack---")
								})
							}
						}
					}
				}
			}
		}
	}
}
