// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysFile is the golang structure of table sys_file for DAO operations like Where/Data.
type SysFile struct {
	g.Meta     `orm:"table:sys_file, do:true"`
	Id         interface{} //
	FilePath   interface{} // 文件目录地址
	FileName   interface{} // 文件名
	FileSizeKb interface{} // 文件大小kb
	FileSuffix interface{} // 文件后缀
	CreatedAt  *gtime.Time // 创建时间
	DeletedAt  *gtime.Time // 删除时间
}