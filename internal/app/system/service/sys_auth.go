package service

import (
	"context"
	"fmt"
	"github.com/Wyatex/SakuraAdminGF/api/v1/system"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/dao"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	insSysAuth = sSysAuth{}
)

type sSysAuth struct{}

func SysAuth() *sSysAuth {
	return &insSysAuth
}

func (s *sSysAuth) Add(ctx context.Context, in *model.SysAuthAddInput) (err error) {
	// 检查name是否已存在
	if err = s.checkAuthIsExist(ctx, in); err != nil {
		return
	}

	// 写入数据
	_, err = dao.SysAuth.Ctx(ctx).Data(in).Insert()
	return
}

// todo：删除操作需要把子权限也删除
func (*sSysAuth) DelById(ctx context.Context, id uint64) (err error) {
	//_, err = dao.SysAuth.Ctx(ctx).Delete(dao.SysAuth.Columns().Id, id)
	return
}

// todo: 需要调用DelById进行删除
func (*sSysAuth) DelByIds(ctx context.Context, ids []uint64) (err error) {
	//_, err = dao.SysAuth.Ctx(ctx).WhereIn(dao.SysAuth.Columns().Id, ids).Delete()
	return
}

func (*sSysAuth) Detail(ctx context.Context, id uint64, res **system.SysAuthDetailRes) (err error) {
	err = dao.SysAuth.Ctx(ctx).Scan(res, dao.SysAuth.Columns().Id, id)
	return
}

func (*sSysAuth) Edit(ctx context.Context, in *entity.SysAuth) (err error) {
	_, err = dao.SysAuth.Ctx(ctx).Data(in).Update()
	return
}

func (*sSysAuth) List(ctx context.Context) (list []*model.SysAuthItem, err error) {
	// list不暂时先不做条件查询
	result, err := dao.SysAuth.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	var rawList []*entity.SysAuth
	err = result.Structs(&rawList)
	if err != nil {
		return nil, err
	}
	list, err = insSysAuth.authListToTree(rawList)
	return
}

// RouteRule 获取路由权限
func (s *sSysAuth) RouteRule(ctx context.Context) (list []*model.SysAuthRoute, home string, err error) {
	uid := ctx.Value("uid").(uint64)
	routeList, err := s.getRouteAuthList(ctx)
	if err != nil {
		return nil, "", err
	}
	// 如果不是超级管理员，则对路由列表进行过滤
	if !SysUser().IsAdmin(uid) {
		var authFilterList []*entity.SysAuth
		userAuthMap := s.getUserAuthMap(uid)
		// 从所有权限列表过滤出有权限的列表
		for _, v := range routeList {
			if _, ok := userAuthMap[v.Id]; ok {
				authFilterList = append(authFilterList, v)
			}
		}
		if len(authFilterList) > 0 {
			routeList = authFilterList
		} else {
			routeList = []*entity.SysAuth{}
		}
	}
	// 取第一个页面为主页
	if len(routeList) > 0 {
		for _, route := range routeList {
			if route.Type == 2 {
				home = route.RouteName
				break
			}
		}
	} else {
		return nil, "", nil
	}

	// 生成树结构路由列表
	list, err = s.authListToRouteTree(routeList)
	if err != nil {
		return nil, "", err
	}
	return
}

// ButtonRule todo:获取按钮权限列表
func (s *sSysAuth) ButtonRule(ctx context.Context) (list []string, err error) {
	uid := ctx.Value("uid").(uint64)
	btnList, err := s.getButtonAuthList(ctx)
	if err != nil {
		return nil, err
	}
	// 如果不是超级管理员，则对按钮权限列表进行过滤
	if !SysUser().IsAdmin(uid) {
		var authFilterList []*entity.SysAuth
		userAuthMap := s.getUserAuthMap(uid)
		// 从所有权限列表过滤出有权限的列表
		for _, v := range btnList {
			if _, ok := userAuthMap[v.Id]; ok {
				authFilterList = append(authFilterList, v)
			}
		}
		if len(authFilterList) > 0 {
			btnList = authFilterList
		} else {
			btnList = []*entity.SysAuth{}
		}
	}
	return
}

// 检查权限是否已存在
func (*sSysAuth) checkAuthIsExist(ctx context.Context, in *model.SysAuthAddInput) (err error) {
	if in.Type == 0 || in.Type == 1 || in.Type == 2 {
		result, err := dao.SysAuth.Ctx(ctx).One(dao.SysAuth.Columns().RouteName, in.RouteName)
		if err != nil {
			return err
		}
		if !result.IsEmpty() {
			return gerror.New("路由名称已存在")
		}
		result, err = dao.SysAuth.Ctx(ctx).One(dao.SysAuth.Columns().RoutePath, in.RoutePath)
		if err != nil {
			return err
		}
		if !result.IsEmpty() {
			return gerror.New("路由路径已存在")
		}
	} else if in.Type == 2 {
		result, err := dao.SysAuth.Ctx(ctx).One(dao.SysAuth.Columns().ApiPath, in.ApiPath)
		if err != nil {
			return err
		}
		if !result.IsEmpty() {
			return gerror.New("api权限已存在")
		}
	} else if in.Type == 3 {
		result, err := dao.SysAuth.Ctx(ctx).One(dao.SysAuth.Columns().BtnName, in.BtnName)
		if err != nil {
			return err
		}
		if !result.IsEmpty() {
			return gerror.New("按钮权限已存在")
		}
	}
	return nil
}

// 权限记录转成树结构
func (*sSysAuth) authListToTree(raw []*entity.SysAuth) (list []*model.SysAuthItem, err error) {
	list = []*model.SysAuthItem{}
	for _, auth := range raw {
		if auth.ParentId == 0 {
			authItem := &model.SysAuthItem{}
			err := gconv.Struct(auth, &authItem)
			if err != nil {
				return nil, err
			}
			err = authItem.SetChildren(raw)
			if err != nil {
				return nil, err
			}
			list = append(list, authItem)
		}
	}
	return
}

// 路由记录转成树结构
func (*sSysAuth) authListToRouteTree(raw []*entity.SysAuth) (list []*model.SysAuthRoute, err error) {
	list = []*model.SysAuthRoute{}
	for _, auth := range raw {
		if auth.ParentId == 0 {
			authItem := &model.SysAuthRoute{
				Path: auth.RoutePath,
				Name: auth.RouteName,
				Meta: model.SysAuthRouteMeta{
					Title: auth.Title,
					Icon:  auth.Icon,
					Order: auth.Sort,
				},
			}
			if auth.Type == 0 {
				authItem.Component = "basic"
			} else if auth.Type == 1 {
				authItem.Component = "multi"
			} else if auth.Type == 2 {
				authItem.Component = "self"
			}
			err = authItem.SetChildren(raw, auth.Id)
			if err != nil {
				return nil, err
			}
			list = append(list, authItem)
		}
	}
	return
}

// 查出所有状态为1路由权限记录
func (*sSysAuth) getRouteAuthList(ctx context.Context) (list []*entity.SysAuth, err error) {
	err = dao.SysAuth.Ctx(ctx).
		WhereIn(dao.SysAuth.Columns().Type, g.Slice{0, 1, 2}).
		Where(dao.SysAuth.Columns().Status, 1).Scan(&list)
	return
}

// 查出所有按钮权限记录
func (*sSysAuth) getButtonAuthList(ctx context.Context) (list []*entity.SysAuth, err error) {
	err = dao.SysAuth.Ctx(ctx).
		Where(dao.SysAuth.Columns().Type, 4).
		Where(dao.SysAuth.Columns().Status, 1).Scan(&list)
	return
}

// 从casbin实例中查出某个用户拥有的权限
func (*sSysAuth) getUserAuthMap(uid uint64) map[uint]struct{} {
	var roleStrIds []string
	// 查出用户的角色列表
	groupPolicy := CasbinEnforcer().GetFilteredGroupingPolicy(0, fmt.Sprintf("%s%d", "u_", uid))
	if len(groupPolicy) > 0 {
		roleStrIds = make([]string, len(groupPolicy))
		for k, v := range groupPolicy {
			roleStrIds[k] = v[1]
		}
	} else {
		roleStrIds = []string{}
	}
	// 遍历出权限列表并转成map
	userAuthMap := map[uint]struct{}{}
	for _, roleId := range roleStrIds {
		pList := CasbinEnforcer().GetFilteredPolicy(0, roleId)
		for _, p := range pList {
			userAuthMap[gconv.Uint(p[1])] = struct{}{}
		}
	}
	return userAuthMap
}
