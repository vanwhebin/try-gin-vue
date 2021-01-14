package middlerware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vanwhebin/try-gin-vue/response"
)

func ErrorRecover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, gin.H{}, fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}
