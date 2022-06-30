package service

import (
	"context"
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/api/v1/system"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/dao"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/Wyatex/SakuraAdminGF/lib/db"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sSysRole struct{}

var insSysRole = sSysRole{}

func SysRole() *sSysRole {
	return &insSysRole
}
func (s *sSysRole) Add(ctx context.Context, in *model.SysRoleAddInput) (err error) {
	// 检查角色是否已存在
	if record, err := dao.SysRole.Ctx(ctx).One(dao.SysRole.Columns().RoleName, in.RoleName); err != nil {
		return err
	} else if !record.IsEmpty() {
		return gerror.Newf(`角色"%s"已存在`, in.RoleName)
	}

	// 写入数据
	_, err = dao.SysRole.Ctx(ctx).Data(in).Insert()
	return
}

func (*sSysRole) DelById(ctx context.Context, id uint64) (err error) {
	_, err = dao.SysRole.Ctx(ctx).Delete(dao.SysRole.Columns().Id, id)
	return
}

func (*sSysRole) DelByIds(ctx context.Context, ids []uint64) (err error) {
	_, err = dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().Id, ids).Delete()
	return
}

func (*sSysRole) Detail(ctx context.Context, id uint64, res **system.SysRoleDetailRes) (err error) {
	err = dao.SysRole.Ctx(ctx).Scan(res, dao.SysRole.Columns().Id, id)
	return
}

func (*sSysRole) Edit(ctx context.Context, in *entity.SysRole) (err error) {
	if record, err := dao.SysRole.Ctx(ctx).WhereNot(dao.SysRole.Columns().Id, in.Id).One(dao.SysRole.Columns().RoleName, in.RoleName); err != nil {
		return err
	} else if !record.IsEmpty() {
		return gerror.Newf(`角色"%s"已存在`, in.RoleName)
	}
	_, err = dao.SysRole.Ctx(ctx).Update(in, dao.SysDept.Columns().Id, in.Id)
	return
}

func (*sSysRole) List(ctx context.Context, param *system.SysRoleListReq) ([]*system.SysRoleListRes, error) {
	m := dao.SysRole.Ctx(ctx).Safe(false)

	// 添加查询条件
	if param.RoleName != "" {
		m.WhereLike(dao.SysRole.Columns().RoleName, db.Like(param.RoleName))
	}
	if param.DefaultRouter != 0 {
		m.Where(dao.SysRole.Columns().DefaultRouter, param.DefaultRouter)
	}
	if param.Remark != "" {
		m.WhereLike(dao.SysRole.Columns().Remark, db.Like(param.Remark))
	}
	if param.StartTime != nil {
		m.WhereGTE(dao.SysRole.Columns().CreatedAt, *param.StartTime)
	}
	if param.EndTime != nil {
		m.WhereLTE(dao.SysRole.Columns().CreatedAt, *param.EndTime)
	}

	// 查询
	var list []*system.SysRoleListRes
	err := m.Scan(&list)
	return list, err
}

func (*sSysRole) GetRoleIds(ctx context.Context) (list []uint, err error) {
	rawList, err := dao.SysRole.Ctx(ctx).Fields("id").Value()
	if err != nil {
		return nil, err
	}
	list = rawList.Uints()
	return
}

func (*sSysRole) Page(ctx context.Context, param *system.SysRolePageReq) (*common.PageRes, error) {
	m := dao.SysRole.Ctx(ctx).Safe(false)

	// 添加查询条件
	if param.RoleName != "" {
		m.WhereLike(dao.SysRole.Columns().RoleName, db.Like(param.RoleName))
	}
	if param.DefaultRouter != 0 {
		m.Where(dao.SysRole.Columns().DefaultRouter, param.DefaultRouter)
	}
	if param.Remark != "" {
		m.WhereLike(dao.SysRole.Columns().Remark, db.Like(param.Remark))
	}
	if param.StartTime != nil {
		m.WhereGTE(dao.SysRole.Columns().CreatedAt, *param.StartTime)
	}
	if param.EndTime != nil {
		m.WhereLTE(dao.SysRole.Columns().CreatedAt, *param.EndTime)
	}

	// 查询
	page := common.PageRes{
		PageNo:   param.PageNo,
		PageSize: param.PageSize,
	}
	var rows []*entity.SysRole
	err := db.PageWithPointer(m, &page, &rows)
	return &page, err
}
