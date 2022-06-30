package db

// Like 两端拼接上百分号
func Like(s string) string {
	return "%" + s + "%"
}

// LikeLeft 左边拼接上百分号
func LikeLeft(s string) string {
	return "%" + s
}

// LikeRight 右边拼接上百分号
func LikeRight(s string) string {
	return s + "%"
}
