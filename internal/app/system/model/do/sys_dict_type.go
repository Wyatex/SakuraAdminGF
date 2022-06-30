// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure of table sys_dict_type for DAO operations like Where/Data.
type SysDictType struct {
	g.Meta    `orm:"table:sys_dict_type, do:true"`
	Id        interface{} //
	DictName  interface{} // 字典中文名
	DictType  interface{} // 字典类型(英文名)
	Status    interface{} // 状态(1:正常,0:关闭)
	Desc      interface{} // 描述
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
