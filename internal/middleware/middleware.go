package middleware

import (
	"github.com/Wyatex/SakuraAdminGF/api/v1/system"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/service"
	"github.com/Wyatex/SakuraAdminGF/internal/consts"
	"github.com/Wyatex/SakuraAdminGF/lib/response"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"net/http"
)

type DefaultResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

// CORS 跨域设置
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// HandleResponse 统一处理响应，仿照ghttp.MiddlewareHandlerResponse设计
func HandleResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			code = gcode.CodeNotFound
		case http.StatusForbidden:
			code = gcode.CodeNotAuthorized
		default:
			code = gcode.CodeUnknown
		}
	} else {
		code = gcode.New(consts.Success, "成功", nil)
		msg = "成功"
	}

	defaultRes := DefaultResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
		Success: code.Code() == consts.Success,
	}

	// 将相应数据放到ctx中，用于日志hook获取
	responseString := gjson.New(defaultRes).MustToJsonString()
	r.SetCtxVar("resJson", responseString)

	// 返回
	r.Response.WriteJson(responseString)

	// 暂时先决定WriteJson后退出路由
	r.Exit()
}

// UserCtx 用于获取SysUser表完整用户信息
func UserCtx(r *ghttp.Request) {
	// 拿到登录的信息
	resp := TokenData(r)
	if !resp.Success() {
		response.Json(r, resp.Code, resp.Msg, nil)
		return
	}
	id := gconv.Uint64(resp.Get("data").Map()["id"])

	// 查询用户并写入Ctx
	var user *system.SysUserDetailRes
	err := service.SysUser().Detail(r.GetCtx(), id, &user)
	if err != nil {
		g.Log().Error(r.GetCtx(), err)
		response.Json(r, consts.Fail, resp.Msg, nil)
		return
	}
	r.SetCtxVar("userDetail", user)
	r.Middleware.Next()
}

func SetModelName(modelName string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		r.SetCtxVar("modelName", modelName)
		r.Middleware.Next()
	}
}
