package middleware

import (
	"github.com/HwGin-original/gin-admin/server/src/ginplus"
	"github.com/HwGin-original/gin-admin/server/pkg/util"
	"github.com/gin-gonic/gin"
)

// TraceMiddleware 跟踪ID中间件
func TraceMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		// 优先从请求头中获取请求ID，如果没有则使用UUID
		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = util.MustUUID()
		}
		c.Set(ginplus.TraceIDKey, traceID)
		c.Next()
	}
}
