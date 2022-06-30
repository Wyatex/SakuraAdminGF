package middleware

import (
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/service"
	"github.com/Wyatex/SakuraAdminGF/internal/consts"
	"github.com/Wyatex/SakuraAdminGF/lib/response"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	ctx          = gctx.New()
	GtokenConfig = gtoken.GfToken{
		CacheMode:        g.Cfg().MustGet(ctx, "gToken.CacheMode").Int8(),
		CacheKey:         g.Cfg().MustGet(ctx, "gToken.CacheKey").String(),
		Timeout:          g.Cfg().MustGet(ctx, "gToken.Timeout").Int(),
		EncryptKey:       g.Cfg().MustGet(ctx, "gToken.EncryptKey").Bytes(),
		MultiLogin:       g.Cfg().MustGet(ctx, "gToken.MultiLogin").Bool(),
		LoginPath:        g.Cfg().MustGet(ctx, "gToken.LoginPath").String(),
		LogoutPath:       g.Cfg().MustGet(ctx, "gToken.LogoutPath").String(),
		AuthExcludePaths: g.Cfg().MustGet(ctx, "gToken.AuthExcludePaths").Strings(),
		AuthAfterFunc:    authAfter,
		LoginBeforeFunc:  loginBefore,
		LoginAfterFunc:   loginAfter,
		MiddlewareType:   gtoken.MiddlewareTypeGroup,
	}
)

// 登陆前校验
func loginBefore(r *ghttp.Request) (string, interface{}) {
	username := r.Get("username").String()
	password := r.Get("password").String()
	user, err := service.SysUser().LoginCheck(r.Context(), username, password)
	if err != nil {
		response.Json(r, -1, err.Error(), nil)
	} else if user == nil {
		response.Json(r, -1, "账号或密码不正确", nil)
	}
	// 生成保存到redis的用户数据
	data := g.Map{
		"id":       user.Id,
		"username": user.Username,
		"nickname": user.Nickname,
		"phone":    user.Phone,
		"email":    user.Email,
		"deptId":   user.DeptId,
		"avatar":   user.Avatar,
	}
	r.SetCtxVar("userData", data)
	r.SetCtxVar("uid", user.Id)
	return username, data
}

// 登录后返回
func loginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		// 查询权限信息
		routes, home, err := service.SysAuth().RouteRule(r.GetCtx())
		if err != nil {
			response.Json(r, consts.Fail, err.Error(), nil)
		}
		// todo: 按钮权限待做
		response.Json(r, consts.Success, "登录成功", g.Map{
			"token":    respData.Get("token"),
			"userInfo": r.GetCtxVar("userData"),
			"route": g.Map{
				"routes": routes,
				"home":   home,
			},
		})
	} else {
		response.Json(r, respData.Code, respData.Msg, nil)
	}
}

// 校验结束后，如果不通过直接结束请求
func authAfter(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Code == gtoken.UNAUTHORIZED {
		response.Json(r, 401, "token已过期或者未携带token", nil)
	} else if respData.Code == gtoken.FAIL {
		response.Json(r, consts.Fail, "校验不通过", nil)
	} else if respData.Code == gtoken.ERROR {
		response.Json(r, respData.Code, respData.Msg, nil)
	}
	userData := respData.Get("data")
	r.SetCtxVar("tokenData", userData)
	r.SetCtxVar("uid", gconv.Uint64(userData.Map()["id"]))
	r.Middleware.Next()
}

func TokenData(r *ghttp.Request) gtoken.Resp {
	return GtokenConfig.GetTokenData(r)
}
