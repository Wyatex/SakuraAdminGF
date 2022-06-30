package service

import (
	"fmt"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

const (
	defaultTableName     = "casbin_rule"
	dropPolicyTableSql   = `DROP TABLE IF EXISTS %s`
	createPolicyTableSql = `
CREATE TABLE IF NOT EXISTS %s (
	ptype VARCHAR(10) NOT NULL DEFAULT '',
	v0 VARCHAR(256) NOT NULL DEFAULT '',
	v1 VARCHAR(256) NOT NULL DEFAULT '',
	v2 VARCHAR(256) NOT NULL DEFAULT ''
) COMMENT = 'policy table';
`
	defaultModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
)

type casbinAdapter struct{}

var (
	enforcer *casbin.SyncedEnforcer
)

// 初始化适配器
func init() {
	insCasbinAdapter := new(casbinAdapter)
	// 从字符串初始化模型
	m, _ := model.NewModelFromString(defaultModel)
	e, err := casbin.NewSyncedEnforcer(m, insCasbinAdapter)
	if err != nil {
		panic(err)
		return
	}
	enforcer = e
	return
}

// CasbinEnforcer 获取Enforcer实例
func CasbinEnforcer() *casbin.SyncedEnforcer {
	return enforcer
}

// Reload 当修改用户、角色、权限时需要重新加载一次
func Reload() error {
	insCasbinAdapter := new(casbinAdapter)
	// 从字符串初始化模型
	m, _ := model.NewModelFromString(defaultModel)
	e, err := casbin.NewSyncedEnforcer(m, insCasbinAdapter)
	if err != nil {
		return err
	}
	enforcer = e
	return nil
}

// 适配器相关方法

// LoadPolicy 从数据库加载策略.
func (a *casbinAdapter) LoadPolicy(model model.Model) (err error) {
	list, err := CasbinRule().All()
	if err != nil {
		return
	}
	for _, rule := range list {
		lineText := rule.PType
		if rule.V0 != "" {
			lineText += ", " + rule.V0
		}
		if rule.V1 != "" {
			lineText += ", " + rule.V1
		}
		if rule.V2 != "" {
			lineText += ", " + rule.V2
		}
		persist.LoadPolicyLine(lineText, model)
	}
	return
}

// SavePolicy 保存策略到数据库.
func (a *casbinAdapter) SavePolicy(model model.Model) (err error) {
	// 直接删除表再创建
	if err = a.dropPolicyTable(); err != nil {
		return
	}
	if err = a.createPolicyTable(); err != nil {
		return
	}

	var list []*entity.CasbinRule
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			list = append(list, convToEntity(ptype, rule))
		}
	}
	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			list = append(list, convToEntity(ptype, rule))
		}
	}
	return CasbinRule().Save(list)
}

// AddPolicy 向存储中添加策略规则
func (a *casbinAdapter) AddPolicy(sec string, ptype string, rule []string) (err error) {
	err = CasbinRule().Add(convToEntity(ptype, rule))
	return
}

// RemoveFilteredPolicy 从存储器中移除可匹配过滤器的策略规则。
func (a *casbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) (err error) {
	line := &entity.CasbinRule{}
	line.PType = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	return
}

// dropPolicyTable 删除表
func (a *casbinAdapter) dropPolicyTable() (err error) {
	_, err = g.DB().Exec(gctx.New(), fmt.Sprintf(dropPolicyTableSql, defaultTableName))
	return
}

// createPolicyTable 创建表
func (a *casbinAdapter) createPolicyTable() (err error) {
	_, err = g.DB().Exec(gctx.New(), fmt.Sprintf(createPolicyTableSql, defaultTableName))
	return
}

// RemovePolicy 从存储中删除策略规则
func (a *casbinAdapter) RemovePolicy(sec string, ptype string, rule []string) (err error) {
	m := convToEntity(ptype, rule)
	err = CasbinRule().Del(g.Map{
		"p_type": ptype,
		"v0": m.V0,
		"v1": m.V1,
		"v2": m.V2,
	})
	return
}

func convToEntity(ptype string, rule []string) *entity.CasbinRule {
	m := &entity.CasbinRule{}
	m.PType = ptype
	if len(rule) > 0 {
		m.V0 = rule[0]
	}
	if len(rule) > 1 {
		m.V1 = rule[1]
	}
	if len(rule) > 2 {
		m.V2 = rule[2]
	}
	return m
}
