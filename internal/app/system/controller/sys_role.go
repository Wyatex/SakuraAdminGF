package controller

import (
	"context"
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/api/v1/system"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	SysRoleController = cSysRole{}
)

type cSysRole struct{}

func (*cSysRole) List(ctx context.Context, req *system.SysRoleListReq) ([]*system.SysRoleListRes, error) {
	list, err := service.SysRole().List(ctx, req)
	return list, err
}

func (*cSysRole) Page(ctx context.Context, req *system.SysRolePageReq) (res *common.PageRes, err error) {
	res, err = service.SysRole().Page(ctx, req)
	return
}

func (*cSysRole) Add(ctx context.Context, req *system.SysRoleAddReq) (res *system.SysRoleAddRes, err error) {
	var param *model.SysRoleAddInput
	err = gconv.Struct(req, &param)
	if err != nil {
		return nil, err
	}
	return nil, service.SysRole().Add(ctx, param)
}

func (*cSysRole) Edit(ctx context.Context, req *system.SysRoleEditReq) (res *system.SysRoleEditRes, err error) {
	var param *entity.SysRole
	err = gconv.Struct(req, &param)
	if err != nil {
		return nil, err
	}
	return nil, service.SysRole().Edit(ctx, param)
}

func (*cSysRole) Del(ctx context.Context, req *system.SysRoleDelReq) (res *system.SysRoleDelRes, err error) {
	if req.Id != 0 {
		return nil, service.SysRole().DelById(ctx, req.Id)
	} else if req.Ids != nil && len(req.Ids) != 0 {
		return nil, service.SysRole().DelByIds(ctx, req.Ids)
	}
	return nil, gerror.New("参数错误")
}

func (*cSysRole) Detail(ctx context.Context, req *system.SysRoleDetailReq) (res *system.SysRoleDetailRes, err error) {
	err = service.SysRole().Detail(ctx, req.Id, &res)
	return
}
