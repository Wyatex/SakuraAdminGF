package model

type SysUserAddInput struct {
	Username string
	Nickname string
	Phone    string
	Address  string
	Email    string
	Avatar   string
	Sex      int
	DeptId   int64
	Remark   string
	Salt     string
	Password string
}

type SysUserEditModel struct {
	Id       uint64
	Username string
	Password string
	Nickname string
	Phone    string
	Address  string
	Email    string
	Avatar   string
	Sex      int
	DeptId   int64
	Remark   string
}
