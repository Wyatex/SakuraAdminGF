// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOpLog is the golang structure for table sys_op_log.
type SysOpLog struct {
	Id           uint64      `json:"id"           ` //
	UserId       uint64      `json:"userId"       ` // 用户id
	Ip           string      `json:"ip"           ` // 请求ip
	Path         string      `json:"path"         ` // 请求路径
	Method       string      `json:"method"       ` // 请求方法
	Status       int         `json:"status"       ` // 是否正常(0正常 1错误)
	Request      string      `json:"request"      ` // 请求body
	Response     string      `json:"response"     ` // 响应Body
	ErrorMessage string      `json:"errorMessage" ` // 错误信息
	CreatedAt    *gtime.Time `json:"createdAt"    ` // 创建时间
}
