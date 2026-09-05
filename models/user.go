package models

// User 用户表
type User struct {
	BaseModel
	Username string `gorm:"column:username;size:64;not null;uniqueIndex:idx_users_username;comment:'用户名'" json:"username"`
	Password string `gorm:"column:password;size:256;not null;comment:'密码(bcrypt)'" json:"-"`
	Nickname string `gorm:"column:nickname;size:128;comment:'昵称'" json:"nickname"`
	Avatar   string `gorm:"column:avatar;size:512;comment:'头像URL'" json:"avatar"`
	Email    string `gorm:"column:email;size:128;comment:'邮箱'" json:"email"`
	Role     string `gorm:"column:role;size:32;default:'admin';comment:'角色'" json:"role"`
}

func (User) TableName() string {
	return "users"
}
