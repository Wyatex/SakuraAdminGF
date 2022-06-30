package service

import (
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/dao"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sCasbinRule struct{}

var (
	insCasbinRule = sCasbinRule{}
	ctx           = gctx.New()
)

func CasbinRule() *sCasbinRule {
	return &insCasbinRule
}

// Save 保存所有的
func (*sCasbinRule) Save(list []*entity.CasbinRule) (err error) {
	_, err = dao.CasbinRule.Ctx(ctx).Data(list).Save()
	return
}

// All 加载所有策略
func (*sCasbinRule) All() (list []*entity.CasbinRule, err error) {
	err = dao.CasbinRule.Ctx(ctx).Scan(&list)
	return
}

// Del 删除一条策略
func (*sCasbinRule) Del(line g.Map) (err error) {
	_, err = dao.CasbinRule.Ctx(ctx).Delete(line)
	return
}

// Add 添加一条策略
func (*sCasbinRule) Add(data *entity.CasbinRule) (err error) {
	_, err = dao.CasbinRule.Ctx(ctx).Insert(data)
	return
}
