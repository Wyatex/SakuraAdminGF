package common

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type PageReq struct {
	PageNo   int `json:"pageNo" v:"required|min:1#页码不能为空|PageNo必须大于等于1"`
	PageSize int `json:"pageSize" v:"required|min:1#每页个数不能为空|PageSize必须大于等于1"`
}

type TimeReq struct {
	StartTime *gtime.Time `json:"startTime"`
	EndTime   *gtime.Time `json:"endTime"`
}

type PageRes struct {
	PageNo    int         `json:"pageNo"`
	PageSize  int         `json:"pageSize"`
	PageTotal int         `json:"pageTotal"`
	Rows      interface{} `json:"rows"`
}
type RecordData struct {
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
	CreatedBy string      `json:"createdBy" ` // 创建人(关联user_id)
	UpdatedBy string      `json:"updatedBy" ` // 更新人(关联user_id)
	DeletedBy string      `json:"deletedBy" ` // 删除人(关联user_id)
}
