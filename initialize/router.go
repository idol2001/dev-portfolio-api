package initialize

import (
	"dev-portfolio-api/api/routes"
	"dev-portfolio-api/middleware"
	"dev-portfolio-api/pkg/global"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/google/uuid"
	sloggin "github.com/samber/slog-gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	// 初始化Trace中间件
	r.Use(func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	})
	// slog日志
	r.Use(sloggin.New(global.Log))
	// 添加全局异常处理中间件
	r.Use(middleware.Exception)
	// GZip压缩插件
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	// 添加跨域中间件, 让请求支持跨域
	// 定义跨域配置
	crosConfig := cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "X-Auth-Token", "X-Real-IP"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	// 注册跨域中间件
	r.Use(cors.New(crosConfig))
	apiGroup := r.Group("dev-portfolio")
	// 方便统一添加路由前缀
	v1Group := apiGroup.Group("v1")
	routes.InitProfileRouter(v1Group)
	return r
}
