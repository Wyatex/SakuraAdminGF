package middleware

import (
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/service"
	"github.com/Wyatex/SakuraAdminGF/lib/response"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func authentication(r *ghttp.Request) {
	userId := r.GetCtxVar("uid").Uint64()
	// 超级管理员直接跳过校验
	if service.SysUser().IsAdmin(userId) {
		r.Middleware.Next()
		return
	}
	groupPolicy := service.CasbinEnforcer().GetFilteredGroupingPolicy(0, gconv.String(userId))
	if len(groupPolicy) == 0 {
		response.FailWithMsg(r, 403, "无权访问")
		return
	}
	// todo: 这里要根据path查询出api的id
	//path := r.Request.URL.Path
	apiId := 0
	canAccess := false
	for _, v := range groupPolicy {
		if service.CasbinEnforcer().HasPolicy(v[1], gconv.String(apiId), "All") {
			canAccess = true
			break
		}
	}
	if !canAccess {
		response.FailWithMsg(r, 403, "无权访问")
		return
	}
	r.Middleware.Next()
}
