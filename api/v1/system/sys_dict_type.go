package system

import (
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/internal/app/common/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

//service
//查询列表返回值
type SysDictTypeListReq struct {
	g.Meta `path:"/sysDictType/list" tags:"SysDictType" method:"post" summary:"分页查询字典"`
	List   []entity.SysDictType `json:"list"`
	Page   int                  `json:"page"`
	Size   int                  `json:"size"`
	Total  int                  `json:"total"`
	common.TimeReq
}

type SysDictTypeListRes entity.SysDictType

//api
//分页请求参数
type SysDictTypePageReq struct {
	g.Meta        `path:"/sysDictType/page" tags:"SysDictType" method:"post" summary:"分页查询字典"`
	DictName      string `p:"dictName"`      //字典名称
	DictType      string `p:"dictType"`      //字典类型
	Status        string `p:"status"`        //字典状态
	BeginTime     string `p:"beginTime"`     //开始时间
	EndTime       string `p:"endTime"`       //结束时间
	OrderByColumn string `p:"orderByColumn"` //排序字段
	IsAsc         string `p:"isAsc"`         //排序方式
	common.PageReq
}
type SysDictTypeRes model.PageOutput

//新增页面请求参数
type SysDictTypeAddReq struct {
	g.Meta   `path:"/sysDictType/add" tags:"SysDictType" method:"post" summary:"添加字典"`
	DictName string `p:"dictName"  v:"required#字典名称不能为空"`
	DictType string `p:"dictType"  v:"required#字典类型不能为空"`
	Status   string `p:"status"  v:"required#状态不能为空"`
	Remark   string `p:"remark"`
}
type SysDictTypeAddRes struct {
	g.Meta `mime:"json" example:"string"`
}

//修改页面请求
type SysDictTypeEditReq struct {
	g.Meta `path:"/sysDictType/edit" tags:"SysDictType" method:"post" summary:"编辑字典"`
	entity.SysDictType
}

type SysDictTypeEditRes struct {
	g.Meta `mime:"json" example:"string"`
}

//删除内容
type SysDictTypeDelReq struct {
	g.Meta `path:"/sysDictType/del" tags:"SysDictType" method:"post" summary:"删除字典"`
	Id     uint64   `json:"id"`
	Ids    []uint64 `json:"ids"`
}

type SysDictTypeDelRes struct {
	g.Meta `mime:"json" example:"string"`
}
