package controller

import (
	"context"
	"github.com/Wyatex/SakuraAdminGF/api/v1/system"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	SysAuthController = cSysAuth{}
)

type cSysAuth struct{}

func (*cSysAuth) List(ctx context.Context, req *system.SysAuthListReq) ([]*system.SysAuthListRes, error) {
	list, err := service.SysAuth().List(ctx)
	if err != nil {
		return nil, err
	}
	var result []*system.SysAuthListRes
	err = gconv.Structs(list, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (*cSysAuth) Add(ctx context.Context, req *system.SysAuthAddReq) (res *system.SysAuthAddRes, err error) {
	var param *model.SysAuthAddInput
	err = gconv.Struct(req, &param)
	if err != nil {
		return nil, err
	}
	return nil, service.SysAuth().Add(ctx, param)
}

func (*cSysAuth) Edit(ctx context.Context, req *system.SysAuthEditReq) (res *system.SysAuthEditRes, err error) {
	var param *entity.SysAuth
	err = gconv.Struct(req, &param)
	if err != nil {
		return nil, err
	}
	return nil, service.SysAuth().Edit(ctx, param)
}

func (*cSysAuth) Del(ctx context.Context, req *system.SysAuthDelReq) (res *system.SysAuthDelRes, err error) {
	if req.Id != 0 {
		return nil, service.SysAuth().DelById(ctx, req.Id)
	} else if req.Ids != nil && len(req.Ids) != 0 {
		return nil, service.SysAuth().DelByIds(ctx, req.Ids)
	}
	return nil, gerror.New("参数错误")
}

func (*cSysAuth) Detail(ctx context.Context, req *system.SysAuthDetailReq) (res *system.SysAuthDetailRes, err error) {
	err = service.SysAuth().Detail(ctx, req.Id, &res)
	return
}

func (*cSysAuth) GetRoute(ctx context.Context, req *system.SysAuthGetRouteReq) (res *system.SysAuthGetRouteRes, err error) {
	// 查询权限信息
	routes, home, err := service.SysAuth().RouteRule(ctx)
	if err != nil {
		return nil, err
	}
	res = &system.SysAuthGetRouteRes{
		Routes: routes,
		Home:   home,
	}
	return
}

//func (*cSysAuth) Rule(ctx context.Context, req *system.SysAuthRuleReq) (res *system.SysAuthRuleRes, err error) {
//	res = &system.SysAuthRuleRes{}
//	routes, err := service.SysAuth().RouteRule(ctx)
//	if err != nil {
//		return nil, err
//	}
//	res.Routes = routes
//	buttons, err := service.SysAuth().ButtonRule(ctx)
//	if err != nil {
//		return nil, err
//	}
//	res.Buttons = buttons
//	return
//}
