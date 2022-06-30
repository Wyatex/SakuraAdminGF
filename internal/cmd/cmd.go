package cmd

import (
	"context"
	. "github.com/Wyatex/SakuraAdminGF/internal/router"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 注册路由
			SetupRouter(ctx, s)
			// 启动HTTP服务器
			s.Run()
			return nil
		},
	}
)
