package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"ytu/ginessential/common"
	"ytu/ginessential/routes"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = routes.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	// 获取当前得工作目录
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")       //设置要读取得文件名
	viper.SetConfigType("yml")               //设置读取得文件类型
	viper.AddConfigPath(workDir + "/config") //设置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
