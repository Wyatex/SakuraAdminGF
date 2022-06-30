package system

import (
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type SysAuthListReq struct {
	g.Meta `path:"/sysAuth/list" tags:"SysAuth" method:"post" summary:"查询权限列表"`
	Name   string `json:"name"`
	common.TimeReq
}

type SysAuthListRes model.SysAuthItem

type SysAuthAddReq struct {
	g.Meta    `path:"/sysAuth/add" tags:"SysAuth" method:"post" summary:"添加权限"`
	ParentId  uint   `json:"parentId"      `                     // 父菜单id
	Type      int    `json:"type" v:"required|min:0|max:3"`      // 类型（0目录 1菜单 2API 3按钮）
	Title     string `json:"menuTitle" v:"required|min:0|max:3"` // 菜单中文展示名
	RoutePath string `json:"routePath"     `                     // 路由目录(唯一)
	RouteName string `json:"routeName"     `                     // 路由name(唯一)
	ApiPath   string `json:"apiPath"       `                     // api路径(唯一)
	BtnName   string `json:"btnName"       `                     // 按钮名称(英文，唯一)
	Icon      string `json:"icon"          `                     // 菜单icon
	Sort      int    `json:"sort"          `                     // 排序
}

type SysAuthAddRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysAuthEditReq struct {
	g.Meta `path:"/sysAuth/edit" tags:"SysAuth" method:"post" summary:"编辑权限"`
	entity.SysAuth
}

type SysAuthEditRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysAuthDelReq struct {
	g.Meta `path:"/sysAuth/del" tags:"SysAuth" method:"post" summary:"删除权限"`
	Id     uint64   `json:"id"`
	Ids    []uint64 `json:"ids"`
}

type SysAuthDelRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysAuthDetailReq struct {
	g.Meta `path:"/sysAuth/detail" tags:"SysAuth" method:"post" summary:"查询权限"`
	Id     uint64 `json:"id" v:"required"`
}

type SysAuthDetailRes struct {
	g.Meta `mime:"json" example:"string"`
	entity.SysAuth
}

type SysAuthRuleReq struct {
	g.Meta `path:"/sysAuth/rule" tags:"SysAuth" method:"post" summary:"查询用户拥有的权限"`
}

type SysAuthRuleRes struct {
	g.Meta  `mime:"json" example:"string"`
	Routes  []*model.SysAuthRoute `json:"routes"`
	Buttons []string              `json:"buttons"`
}

type SysAuthGetRouteReq struct {
	g.Meta `path:"/sysAuth/getRoute" tags:"SysAuth" method:"post" summary:"查询用户拥有权限的路由"`
}

type SysAuthGetRouteRes struct {
	g.Meta `mime:"json" example:"string"`
	Routes []*model.SysAuthRoute `json:"routes"`
	Home   string                `json:"home"`
}
