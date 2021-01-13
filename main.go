package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"vanwhebin/try-gin-vue/common"
)

func main() {
	InitConfig()
	common.InitDB()
	//defer db.Close() // TODO 延时关闭连接
	r := gin.Default()
	r = CollectRoutes(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig(){
	workDir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Load config file error")
	}

}
