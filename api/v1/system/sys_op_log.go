package system

import (
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/Wyatex/SakuraAdminGF/internal/app/common/model"
	"github.com/Wyatex/SakuraAdminGF/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type SysOpLogListReq struct {
	g.Meta    `path:"/sysOpLog/page" tags:"SysOpLog" method:"post" summary:"分页查询日志"`
	UserId    *uint64 `json:"userId"`
	IP        string  `json:"ip"`
	Path      string  `json:"path"`
	Method    string  `json:"method"`
	Status    *int    `json:"status"`
	common.TimeReq
}

type SysOpLogListRes entity.SysOpLog

type SysOpLogPageReq struct {
	g.Meta    `path:"/sysOpLog/page" tags:"SysOpLog" method:"post" summary:"分页查询日志"`
	UserId    *uint64 `json:"userId"`
	IP        string  `json:"ip"`
	Path      string  `json:"path"`
	Method    string  `json:"method"`
	Status    *int    `json:"status"`
	common.TimeReq
	common.PageReq
}

type SysOpLogPageRes model.PageOutput

type SysOpLogDelReq struct {
	g.Meta `path:"/sysOpLog/del" tags:"SysOpLog" method:"post" summary:"删除日志记录"`
	Id     uint64   `json:"id"`
	Ids    []uint64 `json:"ids"`
}

type SysOpLogDelRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysOpLogDetailReq struct {
	g.Meta `path:"/sysOpLog/detail" tags:"SysOpLog" method:"post" summary:"获取一条日志记录"`
	Id     uint64 `json:"id"`
}

type SysOpLogDetailRes struct {
	g.Meta `mime:"json" example:"string"`
}

type SysOpLogClearReq struct {
	g.Meta `path:"/sysOpLog/clear" tags:"SysOpLog" method:"post" summary:"清空所有日志记录"`
}

type SysOpLogClearRes struct {
	g.Meta `mime:"json" example:"string"`
}
