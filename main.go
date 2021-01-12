package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/api/auth/register", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		}

		if len(password) < 6 || len(password) > 18 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位大于18位"})
		}
		// 如果用户名没有传则创建一个随机字符串

		// 用户已存在

		// 创建用户


	})
}


func RandomString (n int) string {
	var letters = "qwertyuiplkjhgfdsazvxbcnmMNHJUYTREFG"


}