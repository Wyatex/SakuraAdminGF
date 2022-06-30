package db

import (
	"github.com/Wyatex/SakuraAdminGF/api/v1/common"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"math"
)

// Page 通用分页查询,返回的rows是结果的map切片类型
func Page(model *gdb.Model, page *common.PageRes) (err error) {
	// 统计总结果数
	count, err := model.Clone().Count()
	if err != nil {
		return err
	}

	// 是否分页查询
	if page.PageNo != 0 && page.PageSize != 0 {
		model = model.Page(page.PageNo, page.PageSize)
	} else {
		return gerror.New("分页页码和每页数量不能为空")
	}

	result, err := model.All()
	page.PageTotal = int(math.Ceil(float64(count) / float64(page.PageSize)))
	page.Rows = result
	return
}

// PageWithPointer 传入带类型的指针,会使用反射将数据转成对应的格式
func PageWithPointer(model *gdb.Model, page *common.PageRes, pointer interface{}) (err error) {
	err = Page(model, page)
	if err != nil {
		return err
	}
	err = page.Rows.(gdb.Result).Structs(pointer)
	if err != nil {
		return err
	}
	page.Rows = pointer
	return
}
