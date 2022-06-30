package controller

import (
	"context"
	"github.com/Wyatex/SakuraAdminGF/api/v1/system"
)

var (
	SysOpLogController = cSysOpLog{}
)

type cSysOpLog struct{}

func (c *cSysOpLog) Page(ctx context.Context, req *system.SysOpLogPageReq) (res *system.SysOpLogPageRes, err error) {
	return
}

func (c *cSysOpLog) Del(ctx context.Context, req *system.SysOpLogDelReq) (res *system.SysOpLogDelRes, err error) {
	return
}

func (c *cSysOpLog) Detail(ctx context.Context, req *system.SysOpLogDetailReq) (res *system.SysOpLogDetailRes, err error) {
	return
}

func (c *cSysOpLog) Clear(ctx context.Context, req *system.SysOpLogClearReq) (res *system.SysOpLogClearRes, err error) {
	return
}
