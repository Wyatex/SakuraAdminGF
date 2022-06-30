package model

import (
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/util/gconv"
)

type SysAuthAddInput struct {
	ParentId  uint
	Type      int
	Title     string
	RoutePath string
	RouteName string
	ApiPath   string
	BtnName   string
	Icon      string
	Sort      int
	Status    int
}

type SysAuthItem struct {
	entity.SysAuth
	Children []*SysAuthItem `json:"children"`
}

type SysAuthRouteMeta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Order int    `json:"order"`
}

type SysAuthRoute struct {
	Meta      SysAuthRouteMeta `json:"meta"`
	Path      string           `json:"path"`
	Name      string           `json:"name"`
	Component string           `json:"component"`
	Children  []*SysAuthRoute  `json:"children"`
}

func (m *SysAuthItem) SetChildren(raw []*entity.SysAuth) error {
	m.Children = []*SysAuthItem{}
	for _, auth := range raw {
		if auth.ParentId == m.Id {
			authItem := &SysAuthItem{}
			err := gconv.Struct(auth, &authItem)
			if err != nil {
				return err
			}
			// 子类还有可能有子类，需要继续递归下去
			err = authItem.SetChildren(raw)
			if err != nil {
				return err
			}
			m.Children = append(m.Children, authItem)
		}
	}
	return nil
}

func (m *SysAuthRoute) SetChildren(raw []*entity.SysAuth, pid uint) error {
	m.Children = []*SysAuthRoute{}
	for _, auth := range raw {
		if auth.ParentId == pid {
			authItem := &SysAuthRoute{
				Path: auth.RoutePath,
				Name: auth.RouteName,
				Meta: SysAuthRouteMeta{
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
			// 子类还有可能有子类，需要继续递归下去
			err := authItem.SetChildren(raw, auth.Id)
			if err != nil {
				return err
			}
			m.Children = append(m.Children, authItem)
		}
	}
	if len(m.Children) == 0 {
		m.Children = nil
	}
	return nil
}
