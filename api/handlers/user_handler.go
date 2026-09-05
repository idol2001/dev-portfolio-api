package handlers

import (
	"dev-portfolio-api/internal/services"
	"dev-portfolio-api/models"
	"strconv"
	"strings"

	cryptoPkg "dev-portfolio-api/pkg/crypto"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers 获取用户列表
func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		models.FailWithMessage("获取用户列表失败", c)
		return
	}
	// 隐藏密码
	for i := range users {
		users[i].Password = ""
	}
	models.OkWithData(users, c)
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}

	// 解密密码（如果已加密）
	password, err := decryptPasswordIfNeeded(req.Password)
	if err != nil {
		models.FailWithMessage("密码解密失败", c)
		return
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		models.FailWithMessage("密码加密失败", c)
		return
	}

	user := &models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Nickname: req.Nickname,
		Email:    req.Email,
		Role:     req.Role,
	}

	if err := services.CreateUserByAdmin(user); err != nil {
		models.FailWithMessage("创建用户失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("创建成功", c)
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}

	var req struct {
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Role     string `json:"role"`
		Avatar   string `json:"avatar"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}

	if err := services.UpdateUser(uint(id), req.Nickname, req.Email, req.Role, req.Avatar); err != nil {
		models.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("更新成功", c)
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}

	var req struct {
		NewPassword string `json:"new_password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}

	// 解密密码（如果已加密）
	password := req.NewPassword
	if strings.HasPrefix(password, "ENC:") {
		encoded := strings.TrimPrefix(password, "ENC:")
		password, err = cryptoPkg.DecryptPassword(encoded)
		if err != nil {
			models.FailWithMessage("密码解密失败", c)
			return
		}
	}

	if err := services.ChangePassword(uint(id), password); err != nil {
		models.FailWithMessage("修改密码失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("密码修改成功", c)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		models.FailWithMessage("参数错误", c)
		return
	}

	// 防止删除最后一个管理员
	if err := services.DeleteUser(uint(id)); err != nil {
		models.FailWithMessage("删除失败："+err.Error(), c)
		return
	}
	models.OkWithMessage("删除成功", c)
}
