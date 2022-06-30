// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/dao/internal"
)

// internalSysDictTypeDao is internal type for wrapping internal DAO implements.
type internalSysDictTypeDao = *internal.SysDictTypeDao

// sysDictTypeDao is the data access object for table sys_dict_type.
// You can define custom methods on it to extend its functionality as you wish.
type sysDictTypeDao struct {
	internalSysDictTypeDao
}

var (
	// SysDictType is globally public accessible object for table sys_dict_type operations.
	SysDictType = sysDictTypeDao{
		internal.NewSysDictTypeDao(),
	}
)

// Fill with you ideas below.