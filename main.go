package main

import (
	"github.com/gin-gonic/gin"
	"vanwhebin/try-gin-vue/common"
)

func main() {
	common.InitDB()
	//defer db.Close() // TODO 延时关闭连接
	r := gin.Default()
	r = CollectRoutes(r)
	panic(r.Run())
}
