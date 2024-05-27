package main

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"gweb/internal/common/sdk/easychat"
	"log"
	"net/http"
)

func init() {
	SetAdapter(gctx.New())
}
func SetAdapter(ctx context.Context) {
	var adapter gcache.Adapter
	adapter = gcache.NewAdapterRedis(g.Redis())
	// 数据库缓存，默认和通用缓冲驱动一致，如果你不想使用默认的，可以自行调整
	g.DB().GetCache().SetAdapter(adapter)
}
func Cors(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// socketio服务
func SocketIo() {
	s := g.Server("socketio")
	s.Use(Cors)
	easy := easychat.New()

	//监听MQ
	go easy.ListenMq()
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	})

	s.BindHandler("/socket.io/", func(r *ghttp.Request) {
		server.ServeHTTP(r.Response.Writer, r.Request)
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		easychat.WriteAllSocketMap(s)
		return nil
	})

	//登录
	server.OnEvent("/", "login", func(s socketio.Conn, msg string) string {
		return easy.Online(s, msg)
	})

	//广播到群
	server.OnEvent("/", "sendRoom", func(s socketio.Conn, msg string) string {
		return easy.SendToGroup(s, msg)
	})

	//退出登录
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		easy.Offline(s)
		defer s.Close()
		return "out"
	})

	//触发未读消息
	server.OnEvent("/", "unRead", func(s socketio.Conn) string {
		return easy.UnRead(s)
	})

	//输入中
	server.OnEvent("/", "entering", func(s socketio.Conn, msg string) string {
		return easy.Entering(s, msg)
	})

	//异常
	server.OnError("/", func(s socketio.Conn, e error) {
		easy.Offline(s)
		g.Log().SetStack(true)
		g.Log().Error(gctx.New(), e)
		log.Println("meet error:", e)
	})

	//断开链接
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		easy.Offline(s)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	s.BindHandler("/debugs", func(r *ghttp.Request) {
		r.Response.WriteJsonExit(g.Map{
			"all":     easychat.OnlineMap(),
			"service": server.Count(),
		})
	})
	s.SetPort(4455)
	s.Run()
}

func main() {
	SocketIo()
}
