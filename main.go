package main

// 导入gin包
import (
	"os"

	"dev-portfolio-api/initialize"
	. "dev-portfolio-api/pkg/global"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// 加载 .env 文件（如果存在），优先级高于配置文件
	_ = godotenv.Load()
	_ = godotenv.Load(".env.local")

	initialize.InitConfig()
	initialize.Logger()
	initialize.InitDatabase()
}

// 入口函数
func main() {
	// 初始化一个http服务对象
	r := initialize.InitRouters()

	// 根据配置设置 Gin 运行模式（支持环境变量覆盖）
	// 优先级：GIN_MODE 环境变量 > 配置文件 system.run-mode > 默认 debug
	ginMode := global.Conf.System.RunMode
	if envMode := os.Getenv("GIN_MODE"); envMode != "" {
		ginMode = envMode
	}
	switch ginMode {
	case "prd", "release", "prod":
		gin.SetMode(gin.ReleaseMode)
	case "st", "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode) // 本地开发默认
	}

	Log.Info("Server is running ...", "mode", gin.Mode())
	r.Run()
}
