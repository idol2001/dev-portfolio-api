package middleware

import (
	"dev-portfolio-api/pkg/global"
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// 全局异常处理中间件
func Exception(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 将异常写入日志
			global.Log.Error(fmt.Sprintf("未知panic异常: %v\n堆栈信息: %v", err, string(debug.Stack())))
			c.Abort()
			return
		}
	}()
	c.Next()
}
