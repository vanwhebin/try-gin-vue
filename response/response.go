package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


 // 封装 统一返回格式
func Response(ctx *gin.Context, httpsStatus int, code uint, data gin.H, msg string) {
	ctx.JSON(httpsStatus, gin.H{"code": code, "data": data, "msg": msg})
}


func Success(ctx * gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}


func Fail(ctx * gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}