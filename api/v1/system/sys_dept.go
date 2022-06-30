package system

import (
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type SysDeptListReq struct {
	g.Meta `path:"/sysDept/list" tags:"SysDept" method:"post" summary:"查询部门列表"`
	Name   string `json:"name"`
	common.TimeReq
}

type SysDeptListRes model.SysDeptItem

type SysDeptAddReq struct {
	g.Meta   `path:"/sysDept/add" tags:"SysDept" method:"post" summary:"添加部门"`
	ParentId uint   `json:"parentId"  `                     // 父级id，0代表根
	DeptName string `json:"deptName"  v:"required#请输入部门名称"` // 部门名称
	Order    int    `json:"order"     `                     // 显示顺序
	LeaderId int64  `json:"leaderId"  `                     // 负责人id
	Status   int    `json:"status"    `                     // 部门状态（1正常，0停用）
}

type SysDeptAddRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysDeptEditReq struct {
	g.Meta `path:"/sysDept/edit" tags:"SysDept" method:"post" summary:"编辑部门"`
	entity.SysDept
}

type SysDeptEditRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysDeptDelReq struct {
	g.Meta `path:"/sysDept/del" tags:"SysDept" method:"post" summary:"删除部门"`
	Id     uint   `json:"id"`
	Ids    []uint `json:"ids"`
}

type SysDeptDelRes struct {
	g.Meta `mime:"json" example:"string"`
}
