package controller

import (
	"context"
	"github.com/Wyatex/SakuraAdminGF/api/v1/system"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	SysDeptController = cSysDept{}
)

type cSysDept struct{}

func (*cSysDept) List(ctx context.Context, req *system.SysDeptListReq) ([]*system.SysDeptListRes, error) {
	list, err := service.SysDept().List(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	var result []*system.SysDeptListRes
	err = gconv.Structs(list, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (*cSysDept) Add(ctx context.Context, req *system.SysDeptAddReq) (res *system.SysDeptAddRes, err error) {
	var param *model.SysDeptAddInput
	err = gconv.Struct(req, &param)
	if err != nil {
		return nil, err
	}
	return nil, service.SysDept().Add(ctx, param)
}

func (*cSysDept) Edit(ctx context.Context, req *system.SysDeptEditReq) (res *system.SysDeptEditRes, err error) {
	var param *entity.SysDept
	err = gconv.Struct(req, &param)
	if err != nil {
		return nil, err
	}
	return nil, service.SysDept().Edit(ctx, param)
}

func (*cSysDept) Del(ctx context.Context, req *system.SysDeptDelReq) (res *system.SysDeptDelRes, err error) {
	//if req.Id != 0 {
	//	return nil, service.SysDept().DelById(ctx, req.Id)
	//} else if req.Ids != nil && len(req.Ids) != 0 {
	//	return nil, service.SysDept().DelByIds(ctx, req.Ids)
	//}
	//return nil, code.ParamError.ToGerror()
	return nil, service.SysDept().Del(ctx, req.Id)
}
