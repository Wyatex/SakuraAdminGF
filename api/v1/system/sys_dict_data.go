package system

import (
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/internal/app/common/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

//service
//查询列表返回值
type SysDictDataListReq struct {
	g.Meta    `path:"/sysDictData/list" tags:"SysDictData" method:"post" summary:"分页查询字典数据"`
	Sort      int         `json:"sort"      ` // 排序标记
	Value     string      `json:"value"     ` // 字典值
	Label     string      `json:"label"     ` // 展示值
	Type      string      `json:"type"      ` // 字典类型
	Status    int         `json:"status"    ` // 状态(1:正常,0:禁用)
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	CreatedBy string      `json:"createdBy" ` // 创建人(关联user_id)
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	UpdatedBy string      `json:"updatedBy" ` // 更新人(关联user_id)
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
	DeletedBy string      `json:"deletedBy" ` // 删除人(关联user_id)
}

type SysDictDataListRes entity.SysDictData

//api
//分页请求参数
type SysDictDataPageReq struct {
	g.Meta   `path:"/sysDictData/page" tags:"SysDictData" method:"post" summary:"分页查询字典数据"`
	DictType string `p:"dictType" v:"required#字典类型不能为空"` //字典类型
	Label    string `json:"label"     `                  // 展示值
	Status   int    `p:"status"`                         //状态
	common.TimeReq
	common.PageReq
}
type SysDictDataRes model.PageOutput

//新增页面请求参数
type SysDictDataAddReq struct {
	g.Meta    `path:"/sysDictData/add" tags:"SysDictData" method:"post" summary:"添加字典数据"`
	DictLabel string `p:"dictLabel"  v:"required#字典标签不能为空"`
	DictValue string `p:"dictValue"  v:"required#字典键值不能为空"`
	DictType  string `p:"dictType"  v:"required#字典类型不能为空"`
	DictSort  int    `p:"dictSort"  v:"required#字典排序不能为空"`
	Status    int    `p:"status"    v:"required#状态不能为空"`
	Remark    string `p:"remark"`
}
type SysDictDataAddRes struct {
	g.Meta `mime:"json" example:"string"`
}

//修改页面请求
type SysDictDataEditReq struct {
	g.Meta `path:"/sysDictData/edit" tags:"SysDictData" method:"post" summary:"编辑字典数据"`
	Id     int64 `p:"dict_data_id" v:"required#主键ID不能为空"`
	entity.SysDictData
}

type SysDictDataEditRes struct {
	g.Meta `mime:"json" example:"string"`
}

//删除内容
type SysDictDataDelReq struct {
	g.Meta `path:"/sysDictData/del" tags:"SysDictData" method:"post" summary:"删除字典数据"`
	Ids    string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

type SysDictDataDelRes struct {
	g.Meta `mime:"json" example:"string"`
}
