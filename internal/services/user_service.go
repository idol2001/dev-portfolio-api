package services

import (
	"dev-portfolio-api/models"
	"dev-portfolio-api/pkg/global"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// GetUsers 获取所有用户
func GetUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	err := global.DB.Order("id ASC").Find(&users).Error
	return users, err
}

// CreateUserByAdmin 后台创建用户
func CreateUserByAdmin(user *models.User) error {
	// 检查用户名是否已存在
	var existing models.User
	err := global.DB.Where("username = ?", user.Username).First(&existing).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	return global.DB.Create(user).Error
}

// UpdateUser 更新用户信息
func UpdateUser(id uint, nickname, email, role, avatar string) error {
	updates := map[string]interface{}{
		"nickname": nickname,
		"email":    email,
		"role":     role,
		"avatar":   avatar,
	}
	return global.DB.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

// ChangePassword 修改用户密码
func ChangePassword(id uint, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	if err != nil {
		return err
	}
	return global.DB.Model(&models.User{}).Where("id = ?", id).Update("password", string(hashedPassword)).Error
}

// DeleteUser 删除用户（保留至少一个管理员）
func DeleteUser(id uint) error {
	var count int64
	global.DB.Model(&models.User{}).Count(&count)
	if count <= 1 {
		return errors.New("至少保留一个管理员账号")
	}
	return global.DB.Delete(&models.User{}, id).Error
}
