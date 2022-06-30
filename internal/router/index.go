package router

import (
	"context"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/controller"
	. "github.com/Wyatex/SakuraAdminGF/internal/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type iRouter interface {
	Setup(group *ghttp.RouterGroup)
}

func SetupRouter(ctx context.Context, s *ghttp.Server) {
	s.Group("/", func(group *ghttp.RouterGroup) {
		// 这里全局注册默认CORS、通用controller响应处理、gtoken登录校验中间件
		// 如果还有需要全局注册的中间件可以加到下面
		group.Middleware(CORS, HandleResponse)
		// 不需要登录校验的路径需要加到gtoken.go里的authExcludePaths
		err := GtokenConfig.Middleware(ctx, group)
		if err != nil {
			g.Log().Error(ctx, err)
			panic("绑定登录校验中间件失败")
			return
		}

		// 操作日志记录到数据库回调
		//group.Hook("/*", ghttp.HookAfterOutput, hook.OperationLog)

		// 注册路由
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Bind(controller.SysDeptController)
			group.Bind(controller.SysUserController)
			group.Bind(controller.SysRoleController)
			group.Bind(controller.SysOpLogController)
			group.Bind(controller.SysAuthController)
		})
	})
	return
}
