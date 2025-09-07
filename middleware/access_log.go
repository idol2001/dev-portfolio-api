package middleware

import (
	"bytes"
	"dev-portfolio-api/pkg/global"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		global.Log.Info(fmt.Sprint("AccessLogWriter Write", err.Error()))
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func (w AccessLogWriter) WriteString(p string) (int, error) {
	if n, err := w.body.WriteString(p); err != nil {
		global.Log.Info(fmt.Sprint("AccessLogWriter WriteString", err.Error()))
		return n, err
	}
	return w.ResponseWriter.WriteString(p)
}
