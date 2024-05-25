package cmd

import (
	"context"
	"gweb/internal/controller/chat"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gweb/internal/service"
)

var (
	// Main is the main command.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of simple gweb demos",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				// Group middlewares.
				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareCORS,
				)
				// Register route handlers.
				var (
					userCtrl = chat.NewUsers()
				)
				group.Bind(
					userCtrl,
				)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
				})
			})

			s.Run()
			return nil
		},
	}
)
