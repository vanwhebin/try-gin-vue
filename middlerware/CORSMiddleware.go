package middlerware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	domain := viper.GetString("client.domain")
	return func(ctx *gin.Context) {
		// 如果是options请求
		// 判断访问来源，判断访问方法
		// 设置返回头部

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", domain)
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		}else {
			ctx.Next()
		}
	}
}
