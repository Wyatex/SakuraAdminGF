package system

import (
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type SysUserListReq struct {
	g.Meta   `path:"/sysUser/list" tags:"SysUser" method:"post" summary:"获取所有用户"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	DeptId   *int64 `json:"deptId"`
	Sex      *int   `json:"sex"`
	common.TimeReq
}

type SysUserListRes entity.SysUser

type SysUserPageReq struct {
	g.Meta   `path:"/sysUser/page" tags:"SysUser" method:"post" summary:"分页获取用户"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	DeptId   *int64 `json:"deptId"`
	Sex      *int   `json:"sex"`
	common.TimeReq
	common.PageReq
}

type SysUserPageRes common.PageRes

type SysUserAddReq struct {
	g.Meta   `path:"/sysUser/add" tags:"SysUser" method:"post" summary:"添加用户"`
	Username string `json:"username" v:"required#请输入账号"`
	Password string `json:"password" v:"required|password#请输入密码|请输入6-18密码"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone" v:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email" v:"email"`
	Avatar   string `json:"avatar"`
	Sex      int    `json:"sex"`
	DeptId   uint   `json:"deptId"`
	Remark   string `json:"remark"`
}

type SysUserAddRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysUserEditReq struct {
	g.Meta   `path:"/sysUser/edit" tags:"SysUser" method:"post" summary:"编辑用户"`
	Id       uint64 `json:"id" v:"required#请传递id"`
	Username string `json:"username" v:"required#请输入账号"`
	Password string `json:"password" v:"required|password#请输入密码|请输入6-18密码"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone" v:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email" v:"email"`
	Avatar   string `json:"avatar"`
	Sex      int    `json:"sex"`
	DeptId   uint   `json:"deptId"`
	Remark   string `json:"remark"`
}

type SysUserEditRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysUserDelReq struct {
	g.Meta `path:"/sysUser/del" tags:"SysUser" method:"post" summary:"删除用户"`
	Id     uint64   `json:"id"`
	Ids    []uint64 `json:"ids"`
}

type SysUserDelRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysUserDetailReq struct {
	g.Meta `path:"/sysUser/detail" tags:"SysUser" method:"post" summary:"查询用户"`
	Id     uint64 `json:"id" v:"required"`
}

type SysUserDetailRes struct {
	g.Meta `mime:"json" example:"string"`
	entity.SysUser
}
