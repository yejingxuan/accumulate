package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// 支持跨域
func Cors() gin.HandlerFunc {
	mwCORS := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type",
			"X-Requested-With", "X-Request-ID", "X-HTTP-Method-Override",
			"Content-Type", "Upload-Length", "Upload-Offset", "Tus-Resumable",
			"Upload-Metadata", "Upload-Defer-Length", "Upload-Concat"},
		ExposeHeaders: []string{"Content-Type", "Upload-Offset", "Location",
			"Upload-Length", "Tus-Version", "Tus-Resumable", "Tus-Max-Size",
			"Tus-Extension", "Upload-Metadata", "Upload-Defer-Length", "Upload-Concat"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 2400 * time.Hour,
	})
	return mwCORS
}
