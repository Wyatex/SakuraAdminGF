package model

type SysOpLogInput struct {
	UserId       uint64
	Ip           string
	Path         string
	Method       string
	Status       int
	Request      string
	Response     string
	ErrorMessage string
}
