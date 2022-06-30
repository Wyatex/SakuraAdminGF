package service

import (
	"context"
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/api/v1/system"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/dao"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/Wyatex/SakuraAdminGF/lib/db"
)

type sSysOpLog struct{}

var insSysOpLog = sSysOpLog{}

func SysOpLog() *sSysOpLog {
	return &insSysOpLog
}

func (*sSysOpLog) Add(ctx context.Context, in *model.SysOpLogInput) (err error) {
	_, err = dao.SysOpLog.Ctx(ctx).Data(in).Insert()
	return
}

func (*sSysOpLog) DelById(ctx context.Context, id uint64) (err error) {
	_, err = dao.SysOpLog.Ctx(ctx).Delete(dao.SysUser.Columns().Id, id)
	return
}

func (*sSysOpLog) DelByIds(ctx context.Context, ids []uint64) (err error) {
	_, err = dao.SysOpLog.Ctx(ctx).WhereIn(dao.SysUser.Columns().Id, ids).Delete()
	return
}

func (*sSysOpLog) Detail(ctx context.Context, id uint64) (sysOpLog *entity.SysOpLog, err error) {
	err = dao.SysUser.Ctx(ctx).Scan(&sysOpLog, dao.SysUser.Columns().Id, id)
	return
}

func (*sSysOpLog) List(ctx context.Context, param *system.SysOpLogListReq) (list []*system.SysOpLogListRes, err error) {
	m := dao.SysUser.Ctx(ctx).Safe(false)

	// 添加查询条件
	if param.IP != "" {
		m.WhereLike(dao.SysOpLog.Columns().Ip, db.Like(param.IP))
	}
	if param.Path != "" {
		m.WhereLike(dao.SysOpLog.Columns().Path, db.Like(param.Path))
	}
	if param.Method != "" {
		m.WhereLike(dao.SysOpLog.Columns().Method, db.Like(param.Method))
	}
	if param.Status != nil {
		m.Where(dao.SysOpLog.Columns().Status, *param.Status)
	}
	if param.UserId != nil {
		m.Where(dao.SysOpLog.Columns().UserId, *param.UserId)
	}
	if param.StartTime != nil {
		m.WhereGTE(dao.SysOpLog.Columns().CreatedAt, *param.StartTime)
	}
	if param.EndTime != nil {
		m.WhereLTE(dao.SysOpLog.Columns().CreatedAt, *param.EndTime)
	}

	// 查询
	err = m.Scan(&list)
	return list, err
}

func (*sSysOpLog) Page(ctx context.Context, param *system.SysOpLogPageReq) (*common.PageRes, error) {
	m := dao.SysUser.Ctx(ctx).Safe(false)

	// 添加查询条件
	if param.IP != "" {
		m.WhereLike(dao.SysOpLog.Columns().Ip, db.Like(param.IP))
	}
	if param.Path != "" {
		m.WhereLike(dao.SysOpLog.Columns().Path, db.Like(param.Path))
	}
	if param.Method != "" {
		m.WhereLike(dao.SysOpLog.Columns().Method, db.Like(param.Method))
	}
	if param.Status != nil {
		m.Where(dao.SysOpLog.Columns().Status, *param.Status)
	}
	if param.UserId != nil {
		m.Where(dao.SysOpLog.Columns().UserId, *param.UserId)
	}
	if param.StartTime != nil {
		m.WhereGTE(dao.SysOpLog.Columns().CreatedAt, *param.StartTime)
	}
	if param.EndTime != nil {
		m.WhereLTE(dao.SysOpLog.Columns().CreatedAt, *param.EndTime)
	}

	// 查询
	page := common.PageRes{}
	var rows []*entity.SysOpLog
	err := db.PageWithPointer(m, &page, &rows)
	return &page, err
}
