package model

type PageInput struct {
	PageNo   int
	PageSize int
}

type PageOutput struct {
	PageNo    int
	PageSize  int
	PageTotal int
	Rows      interface{}
}

func (page *PageOutput) SetPage(newPage *PageOutput) {
	page.PageNo = newPage.PageNo
	page.PageSize = newPage.PageSize
	page.PageTotal = newPage.PageTotal
	page.Rows = newPage.Rows
}
