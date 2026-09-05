package handlers

import (
	"dev-portfolio-api/internal/services"
	"dev-portfolio-api/middleware"
	"dev-portfolio-api/models"
	cryptoPkg "dev-portfolio-api/pkg/crypto"
	"golang.org/x/crypto/bcrypt"
	"strings"

	"github.com/gin-gonic/gin"
)

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Encrypted bool   `json:"encrypted"` // 标记密码是否已加密
}

// GetPublicKey 返回 RSA 公钥（供前端加密密码使用）
func GetPublicKey(c *gin.Context) {
	pubKey, err := cryptoPkg.GetPublicKeyPEM()
	if err != nil {
		models.FailWithMessage("获取公钥失败: "+err.Error(), c)
		return
	}
	models.OkWithData(gin.H{"publicKey": pubKey}, c)
}

// decryptPasswordIfNeeded 如果密码是加密的，进行解密
func decryptPasswordIfNeeded(encryptedPassword string) (string, error) {
	if encryptedPassword == "" {
		return "", nil
	}
	// 如果密码以 "ENC:" 前缀开头，说明是加密后的密码
	if strings.HasPrefix(encryptedPassword, "ENC:") {
		encoded := strings.TrimPrefix(encryptedPassword, "ENC:")
		return cryptoPkg.DecryptPassword(encoded)
	}
	// 否则直接返回明文密码（向后兼容）
	return encryptedPassword, nil
}

// Login 用户登录
func Login(c *gin.Context) {
	var req LoginRequest
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

	// 验证用户名密码
	user, err := services.Authenticate(req.Username, password)
	if err != nil {
		models.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 生成 Token
	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		models.FailWithMessage("生成令牌失败", c)
		return
	}

	models.OkWithData(gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	}, c)
}

// Register 用户注册
func Register(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.FailWithMessage("参数错误："+err.Error(), c)
		return
	}

	// 创建用户
	user, err := services.CreateUser(req.Username, req.Password)
	if err != nil {
		models.FailWithMessage("创建用户失败："+err.Error(), c)
		return
	}

	// 生成 Token
	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		models.FailWithMessage("生成令牌失败", c)
		return
	}

	models.OkWithData(gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	}, c)
}

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		models.FailWithMessage("未找到用户信息", c)
		return
	}

	user, err := services.GetUserByID(userId.(uint))
	if err != nil {
		models.FailWithMessage("获取用户信息失败", c)
		return
	}

	models.OkWithData(gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"avatar":   user.Avatar,
		"role":     user.Role,
	}, c)
}

// 密码加密
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// 密码验证
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
