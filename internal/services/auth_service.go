package services

import (
	"dev-portfolio-api/models"
	"dev-portfolio-api/pkg/global"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate 验证用户登录
func Authenticate(username, password string) (*models.User, error) {
	user := &models.User{}
	err := global.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	return user, nil
}

// CreateUser 创建用户
func CreateUser(username, password string) (*models.User, error) {
	// 检查用户名是否存在
	var existing models.User
	err := global.DB.Where("username = ?", username).First(&existing).Error
	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	err = global.DB.Create(user).Error
	if err != nil {
		return nil, err
	}

	// 创建默认个人资料
	profileInfo := &models.ProfileInfo{
		Name: user.Username,
	}
	global.DB.Create(profileInfo)

	return user, nil
}

// GetUserByID 根据 ID 获取用户
func GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}
	err := global.DB.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := global.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
