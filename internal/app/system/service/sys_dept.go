package service

import (
	"context"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/dao"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/Wyatex/SakuraAdminGF/lib/db"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysDept struct{}

var (
	insSysDept = sSysDept{}
)

func SysDept() *sSysDept {
	return &insSysDept
}

func (s *sSysDept) List(ctx context.Context, name string) (list []*model.SysDeptItem, err error) {
	if name != "" {
		// 根据名字查的就不需要树结构了
		err = dao.SysDept.Ctx(ctx).WhereLike(dao.SysDept.Columns().DeptName, db.Like(name)).Scan(&list)
		return
	}
	var rawList []*entity.SysDept
	err = dao.SysDept.Ctx(ctx).Scan(&rawList)
	if err != nil {
		return
	}
	list, err = s.deptListToTree(rawList)
	return
}

func (s *sSysDept) Add(ctx context.Context, in *model.SysDeptAddInput) (err error) {
	result, err := dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().ParentId, in.ParentId).Where(dao.SysDept.Columns().DeptName, in.DeptName).Limit(1).All()
	if err != nil {
		return err
	} else if result.Len() > 0 {
		return gerror.Newf(`部门"%s"已存在`, in.DeptName)
	}
	_, err = dao.SysDept.Ctx(ctx).Data(in).Insert()
	return
}

func (s *sSysDept) Edit(ctx context.Context, m *entity.SysDept) (err error) {
	result, err := dao.SysDept.Ctx(ctx).
		Where(dao.SysDept.Columns().ParentId, m.ParentId).
		Where(dao.SysDept.Columns().DeptName, m.DeptName).
		WhereNot(dao.SysDept.Columns().Id, m.Id).
		Limit(1).All()
	if err != nil {
		return err
	} else if result.Len() > 0 {
		return gerror.Newf(`部门"%s"已存在`, m.DeptName)
	}
	_, err = dao.SysDept.Ctx(ctx).Update(m, dao.SysDept.Columns().Id, m.Id)
	return
}

func (s *sSysDept) Del(ctx context.Context, id uint) (err error) {
	_, err = dao.SysDept.Ctx(ctx).Delete(dao.SysDept.Columns().Id, id)
	return
}

// 权限记录转成树结构
func (*sSysDept) deptListToTree(raw []*entity.SysDept) (list []*model.SysDeptItem, err error) {
	list = []*model.SysDeptItem{}
	for _, dept := range raw {
		if dept.ParentId == 0 {
			deptItem := &model.SysDeptItem{}
			err := gconv.Struct(dept, &deptItem)
			if err != nil {
				return nil, err
			}
			err = deptItem.SetChildren(raw)
			if err != nil {
				return nil, err
			}
			list = append(list, deptItem)
		}
	}
	return
}
