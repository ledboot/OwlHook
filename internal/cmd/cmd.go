package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/ledboot/OwlHook/internal/controller/webhook"
	"github.com/ledboot/OwlHook/internal/logic/notifiers"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			notifiers.Init(ctx)

			root := s.Group("/")
			root.Middleware(ghttp.MiddlewareHandlerResponse)
			root.Bind(webhook.NewV1())
			s.Run()
			return nil
		},
	}
)
