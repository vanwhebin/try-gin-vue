package middlerware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vanwhebin/try-gin-vue/common"
	"vanwhebin/try-gin-vue/model"
)

func CheckJwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context){
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Unauthorized"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claim, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Unauthorized"})
			ctx.Abort()
			return
		}

		// 判断用户信息
		userID := claim.UserID
		DB:= common.DB
		var user model.User
		DB.First(&user, userID)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Unauthorized"})
			ctx.Abort()
			return
		}

		// 将用户信息写入情景上下文
		ctx.Set("user", user)
		ctx.Next()

	}
}
