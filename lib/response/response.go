package response

import (
	"github.com/Wyatex/SakuraAdminGF/internal/consts"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Response struct {
	// 代码
	Code int `json:"code"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Message string `json:"message"`
	// 是否成功
	Success bool `json:"success"`
}

// Json 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data interface{}) {
	r.Response.WriteJson(Response{
		Code:    code,
		Message: message,
		Data:    data,
		Success: code == consts.Success,
	})
	r.Exit()
}

func Success(r *ghttp.Request) {
	Json(r, consts.Success, "成功", nil)
}

func SuccessWithData(r *ghttp.Request, data interface{}) {
	Json(r, consts.Success, "成功", data)
}

func Fail(r *ghttp.Request) {
	Json(r, consts.Fail, "操作失败", nil)
}

func FailWithMsg(r *ghttp.Request, code int, msg string) {
	Json(r, code, msg, nil)
}
