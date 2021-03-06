// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/dao/internal"
)

// internalSysOpLogDao is internal type for wrapping internal DAO implements.
type internalSysOpLogDao = *internal.SysOpLogDao

// sysOpLogDao is the data access object for table sys_op_log.
// You can define custom methods on it to extend its functionality as you wish.
type sysOpLogDao struct {
	internalSysOpLogDao
}

var (
	// SysOpLog is globally public accessible object for table sys_op_log operations.
	SysOpLog = sysOpLogDao{
		internal.NewSysOpLogDao(),
	}
)

// Fill with you ideas below.
