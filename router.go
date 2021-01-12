package main

import (
"github.com/gin-gonic/gin"
"vanwhebin/try-gin-vue/controller"
	"vanwhebin/try-gin-vue/middlerware"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middlerware.CheckJwtToken(), controller.Info)
	return r
}
