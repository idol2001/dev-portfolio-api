package main

// 导入gin包
import (
	"context"
	"dev-portfolio-api/initialize"
	. "dev-portfolio-api/pkg/global"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	host := "0.0.0.0"
	port := Conf.System.Port
	// 服务启动及优雅关闭
	// 参考地址https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Log.Error(fmt.Sprint("listen error: ", err))
		}
	}()
	Log.Info(fmt.Sprintf("HTTP Server is running at %s:%d/%s", host, port, Conf.System.UrlPathPrefix))
	// certManager := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist(Conf.System.BaseApi),
	// 	Cache:      autocert.DirCache("./cache"),
	// }
	// Log.Info(fmt.Sprintf("HTTPS Server is running at %s:%d/%s", host, 443, Conf.System.UrlPathPrefix))
	// autotls.RunWithManager(r, &certManager)
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	Log.Info("Shutting down server...")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		Log.Error(fmt.Sprint("Server forced to shutdown: ", err))
	}
	Log.Info("Server exiting")
}
