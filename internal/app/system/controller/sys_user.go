package controller

import (
	"context"
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/api/v1/system"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	SysUserController = cSysUser{}
)

type cSysUser struct{}

func (*cSysUser) List(ctx context.Context, req *system.SysUserListReq) ([]*system.SysUserListRes, error) {
	list, err := service.SysUser().List(ctx, req)
	return list, err
}

func (*cSysUser) Page(ctx context.Context, req *system.SysUserPageReq) (res *common.PageRes, err error) {
	res, err = service.SysUser().Page(ctx, req)
	return
}

func (*cSysUser) Add(ctx context.Context, req *system.SysUserAddReq) (res *system.SysUserAddRes, err error) {
	var param *model.SysUserAddInput
	err = gconv.Struct(req, &param)
	if err != nil {
		return nil, err
	}
	return nil, service.SysUser().Add(ctx, param)
}

func (*cSysUser) Edit(ctx context.Context, req *system.SysUserEditReq) (res *system.SysUserEditRes, err error) {
	var param *model.SysUserEditModel
	err = gconv.Struct(req, &param)
	if err != nil {
		return nil, err
	}
	return nil, service.SysUser().Edit(ctx, param)
}

func (*cSysUser) Del(ctx context.Context, req *system.SysUserDelReq) (res *system.SysUserDelRes, err error) {
	if req.Id != 0 {
		return nil, service.SysUser().DelById(ctx, req.Id)
	} else if req.Ids != nil && len(req.Ids) != 0 {
		return nil, service.SysUser().DelByIds(ctx, req.Ids)
	}
	return nil, gerror.New("参数错误")
}

func (*cSysUser) Detail(ctx context.Context, req *system.SysUserDetailReq) (res *system.SysUserDetailRes, err error) {
	err = service.SysUser().Detail(ctx, req.Id, &res)
	return
}
