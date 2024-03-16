package main

import (
	"demo/config"
	"demo/database"
	"demo/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	// "fmt"
)

func main() {
	config.ConfigInit()
	database.InitMysql()
	database.InitRedis()
	// fmt.Println("Hello World!")
	r := gin.Default()

	r = routers.CollectRouter(r)

	// 从viper中获取到运行端口
	port := viper.GetString("server.port")
	// 如果指定了端口
	if port != "" {
		panic(r.Run(":" + port))
	}
	//	没指定端口就直接运行
	panic(r.Run())

}
