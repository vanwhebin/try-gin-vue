package main

import (
"github.com/gin-gonic/gin"
"vanwhebin/try-gin-vue/controller"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}