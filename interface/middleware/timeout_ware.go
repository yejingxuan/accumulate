package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"op-register/infrastructure/config"
	"time"
)

//gin超时时间设置
func TimeoutMiddleware() func(c *gin.Context) {
	timeout := time.Second * time.Duration(config.CoreConf.Server.MaxHTTPTime)
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}
			cancel()
		}()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
