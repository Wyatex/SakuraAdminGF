package service

import (
	"context"
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/api/v1/system"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/dao"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/Wyatex/SakuraAdminGF/lib/db"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type sSysUser struct{}

var (
	insSysUser     = sSysUser{}
	superAdminList = g.Cfg().MustGet(ctx, "authentication.superAdminIds").Uint64s()
)

func SysUser() *sSysUser {
	return &insSysUser
}

func (s *sSysUser) Add(ctx context.Context, in *model.SysUserAddInput) (err error) {
	dao.SysUser.Ctx(ctx)
	// 检查用户名是否已存在
	if record, err := dao.SysUser.Ctx(ctx).One(dao.SysUser.Columns().Username, in.Username); err != nil {
		return err
	} else if !record.IsEmpty() {
		return gerror.Newf(`账号"%s"已被占用`, in.Username)
	}
	if record, err := dao.SysUser.Ctx(ctx).One(dao.SysUser.Columns().Phone, in.Phone); err != nil {
		return err
	} else if !record.IsEmpty() {
		return gerror.Newf(`手机号"%s"已被占用`, in.Phone)
	}

	dao.SysUser.Ctx(ctx)
	// 登录密码加密
	salt := grand.S(8)
	in.Password = gmd5.MustEncryptString(gmd5.MustEncryptString(in.Password) + gmd5.MustEncryptString(salt))
	in.Salt = salt

	// 写入数据
	_, err = dao.SysUser.Ctx(ctx).Data(in).Insert()
	return
}

func (*sSysUser) DelById(ctx context.Context, id uint64) (err error) {
	_, err = dao.SysUser.Ctx(ctx).Delete(dao.SysUser.Columns().Id, id)
	return
}

func (*sSysUser) DelByIds(ctx context.Context, ids []uint64) (err error) {
	_, err = dao.SysUser.Ctx(ctx).WhereIn(dao.SysUser.Columns().Id, ids).Delete()
	return
}

func (*sSysUser) Detail(ctx context.Context, id uint64, res **system.SysUserDetailRes) (err error) {
	err = dao.SysUser.Ctx(ctx).Scan(res, dao.SysUser.Columns().Id, id)
	return
}

func (*sSysUser) Edit(ctx context.Context, in *model.SysUserEditModel) (err error) {
	// 检查用户名和手机号是否已存在
	if record, err := dao.SysUser.Ctx(ctx).WhereNot(dao.SysUser.Columns().Id, in.Id).One(dao.SysUser.Columns().Username, in.Username); err != nil {
		return err
	} else if !record.IsEmpty() {
		return gerror.Newf(`账号"%s"已被占用`, in.Username)
	}
	if record, err := dao.SysUser.Ctx(ctx).WhereNot(dao.SysUser.Columns().Id, in.Id).One(dao.SysUser.Columns().Phone, in.Phone); err != nil {
		return err
	} else if !record.IsEmpty() {
		return gerror.Newf(`手机号"%s"已被占用`, in.Phone)
	}
	_, err = dao.SysUser.Ctx(ctx).Update(in, dao.SysUser.Columns().Id, in.Id)
	return
}

func (*sSysUser) List(ctx context.Context, param *system.SysUserListReq) ([]*system.SysUserListRes, error) {
	m := dao.SysUser.Ctx(ctx).Safe(false).FieldsEx(dao.SysUser.Columns().Password, dao.SysUser.Columns().Salt)

	// 添加查询条件
	if param.Username != "" {
		m.WhereLike(dao.SysUser.Columns().Username, db.Like(param.Username))
	}
	if param.Nickname != "" {
		m.WhereLike(dao.SysUser.Columns().Nickname, db.Like(param.Nickname))
	}
	if param.Address != "" {
		m.WhereLike(dao.SysUser.Columns().Address, db.Like(param.Address))
	}
	if param.Email != "" {
		m.WhereLike(dao.SysUser.Columns().Email, db.Like(param.Email))
	}
	if param.Phone != "" {
		m.WhereLike(dao.SysUser.Columns().Phone, db.Like(param.Phone))
	}
	if param.Sex != nil {
		m.Where(dao.SysUser.Columns().Sex, *param.Sex)
	}
	if param.DeptId != nil {
		m.Where(dao.SysUser.Columns().DeptId, *param.DeptId)
	}
	if param.StartTime != nil {
		m.WhereGTE(dao.SysUser.Columns().CreatedAt, *param.StartTime)
	}
	if param.EndTime != nil {
		m.WhereLTE(dao.SysUser.Columns().CreatedAt, *param.EndTime)
	}

	// 查询
	var list []*system.SysUserListRes
	err := m.Scan(&list)
	return list, err
}

func (*sSysUser) Page(ctx context.Context, param *system.SysUserPageReq) (*common.PageRes, error) {
	m := dao.SysUser.Ctx(ctx).Safe(false).FieldsEx(dao.SysUser.Columns().Password, dao.SysUser.Columns().Salt)

	// 添加查询条件
	if param.Username != "" {
		m.WhereLike(dao.SysUser.Columns().Username, db.Like(param.Username))
	}
	if param.Nickname != "" {
		m.WhereLike(dao.SysUser.Columns().Nickname, db.Like(param.Nickname))
	}
	if param.Address != "" {
		m.WhereLike(dao.SysUser.Columns().Address, db.Like(param.Address))
	}
	if param.Email != "" {
		m.WhereLike(dao.SysUser.Columns().Email, db.Like(param.Email))
	}
	if param.Phone != "" {
		m.WhereLike(dao.SysUser.Columns().Phone, db.Like(param.Phone))
	}
	if param.Sex != nil {
		m.Where(dao.SysUser.Columns().Sex, *param.Sex)
	}
	if param.DeptId != nil {
		m.Where(dao.SysUser.Columns().DeptId, *param.DeptId)
	}
	if param.StartTime != nil {
		m.WhereGTE(dao.SysUser.Columns().CreatedAt, *param.StartTime)
	}
	if param.EndTime != nil {
		m.WhereLTE(dao.SysUser.Columns().CreatedAt, *param.EndTime)
	}

	// 查询
	page := common.PageRes{
		PageNo:   param.PageNo,
		PageSize: param.PageSize,
	}
	var rows []*entity.SysUser
	err := db.PageWithPointer(m, &page, &rows)
	return &page, err
}

// LoginCheck 检查账号密码是否正确，如果正确返回用户信息
func (*sSysUser) LoginCheck(ctx context.Context, username string, password string) (*entity.SysUser, error) {
	var user *entity.SysUser
	err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Username, username).Limit(1).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user != nil && user.Password != gmd5.MustEncryptString(gmd5.MustEncryptString(password)+gmd5.MustEncryptString(user.Salt)) {
		return nil, nil
	}
	return user, nil
}

func (*sSysUser) IsAdmin(id uint64) bool {
	for _, item := range superAdminList {
		if item == id {
			return true
		}
	}
	return false
}
