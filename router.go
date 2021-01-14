package main

import (
	"github.com/gin-gonic/gin"
	"vanwhebin/try-gin-vue/controller"
	"vanwhebin/try-gin-vue/middlerware"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.Use(middlerware.CORSMiddleware(), middlerware.ErrorRecover())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middlerware.CheckJwtToken(), controller.Info)
	r.GET("/api/auth/logout", middlerware.CheckJwtToken(), controller.Logout)

	categoryRoutes := r.Group("/api/categories")
	categoryController := controller.NewCategory()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT(":id", categoryController.Update)
	categoryRoutes.GET(":id", categoryController.Show)
	categoryRoutes.DELETE(":id", categoryController.Delete)

	postRoutes := r.Group("/api/posts")
	postController := controller.NewPostController()
	postRoutes.POST("list", postController.PageList)
	postRoutes.POST("", middlerware.CheckJwtToken(), postController.Create)
	postRoutes.PUT(":id", middlerware.CheckJwtToken(), postController.Update)
	postRoutes.GET(":id", middlerware.CheckJwtToken(), postController.Show)
	postRoutes.DELETE(":id", middlerware.CheckJwtToken(), postController.Delete)

	return r
}
