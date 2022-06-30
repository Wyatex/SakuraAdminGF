// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/dao/internal"
)

// internalSysAuthDao is internal type for wrapping internal DAO implements.
type internalSysAuthDao = *internal.SysAuthDao

// sysAuthDao is the data access object for table sys_auth.
// You can define custom methods on it to extend its functionality as you wish.
type sysAuthDao struct {
	internalSysAuthDao
}

var (
	// SysAuth is globally public accessible object for table sys_auth operations.
	SysAuth = sysAuthDao{
		internal.NewSysAuthDao(),
	}
)

// Fill with you ideas below.
