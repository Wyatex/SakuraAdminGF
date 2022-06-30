package system

import (
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type SysRoleListReq struct {
	g.Meta        `path:"/sysRole/list" tags:"SysUser" method:"post" summary:"获取所有角色"`
	RoleName      string `json:"roleName"          ` // 角色名
	DefaultRouter uint   `json:"defaultRouterName" ` // 默认打开菜单
	Remark        string `json:"remark"            ` // 备注
	common.TimeReq
}
type SysRoleListRes entity.SysRole

type SysRolePageReq struct {
	g.Meta        `path:"/sysRole/page" tags:"SysRole" method:"post" summary:"分页获取角色"`
	RoleName      string `json:"roleName"          ` // 角色名
	DefaultRouter uint   `json:"defaultRouterName" ` // 默认打开菜单
	Remark        string `json:"remark"            ` // 备注
	common.TimeReq
	common.PageReq
}

type SysRolePageRes common.PageRes

type SysRoleAddReq struct {
	g.Meta            `path:"/sysRole/add" tags:"SysRole" method:"post" summary:"添加角色"`
	RoleName          string `json:"roleName" v:"required#请输入角色名"`
	DefaultRouterName string `json:"defaultRouterName"`
	Remark            string
}

type SysRoleAddRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysRoleEditReq struct {
	g.Meta `path:"/sysRole/edit" tags:"SysRole" method:"post" summary:"编辑角色"`
	entity.SysRole
}

type SysRoleEditRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysRoleDelReq struct {
	g.Meta `path:"/sysRole/del" tags:"SysRole" method:"post" summary:"删除角色"`
	Id     uint64   `json:"id"`
	Ids    []uint64 `json:"ids"`
}

type SysRoleDelRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysRoleDetailReq struct {
	g.Meta `path:"/sysRole/detail" tags:"SysRole" method:"post" summary:"查询角色"`
	Id     uint64 `json:"id" v:"required"`
}

type SysRoleDetailRes struct {
	g.Meta `mime:"json" example:"string"`
	entity.SysRole
}
