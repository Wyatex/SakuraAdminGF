package model

import (
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/util/gconv"
)

type SysDeptItem struct {
	entity.SysDept
	Children []*SysDeptItem `json:"children"`
}

type SysDeptAddInput struct {
	ParentId uint
	DeptName string
	Order    int
	LeaderId int64
	Status   int
}

func (m *SysDeptItem) SetChildren(raw []*entity.SysDept) error {
	m.Children = []*SysDeptItem{}
	for _, dept := range raw {
		if dept.ParentId == m.Id {
			deptItem := &SysDeptItem{}
			err := gconv.Struct(dept, &deptItem)
			if err != nil {
				return err
			}
			// 子类还有可能有子类，需要继续递归下去
			err = deptItem.SetChildren(raw)
			if err != nil {
				return err
			}
			m.Children = append(m.Children, deptItem)
		}
	}
	if len(m.Children) == 0 {
		m.Children = nil
	}
	return nil
}
