package main

// 导入gin包
import (
	"dev-portfolio-api/initialize"
	. "dev-portfolio-api/pkg/global"

	"github.com/gin-gonic/gin"
)

func init() {
	initialize.InitConfig()
	initialize.Logger()
	initialize.Mysql()
}

// 入口函数
func main() {
	// 初始化一个http服务对象
	r := initialize.InitRouters()
	gin.SetMode(gin.DebugMode)
	// host := "0.0.0.0"
	// port := 8080
	Log.Info("Server is running ...")
	r.Run()
}
