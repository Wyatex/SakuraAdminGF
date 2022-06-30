// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAuth is the golang structure for table sys_auth.
type SysAuth struct {
	Id        uint        `json:"id"        ` //
	ParentId  uint        `json:"parentId"  ` // 父菜单id
	Type      int         `json:"type"      ` // 类型（0目录 1多层目录中间层 2页面 3API 4按钮）
	Title     string      `json:"title"     ` // 中文展示名
	RoutePath string      `json:"routePath" ` // 路由目录(唯一)
	RouteName string      `json:"routeName" ` // 路由name(唯一)
	ApiPath   string      `json:"apiPath"   ` // api路径(唯一)
	BtnName   string      `json:"btnName"   ` // 按钮名称(英文，唯一)
	Icon      string      `json:"icon"      ` // 菜单icon
	Sort      int         `json:"sort"      ` // 排序
	Hidden    int         `json:"hidden"    ` // 是否在菜单隐藏（0否 1是）
	Status    int         `json:"status"    ` // 状态(1可用 0不可用)
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
}