package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"ytu/ginessential/common"
	"ytu/ginessential/config"
	"ytu/ginessential/routes"
)

func main() {
	config.InitConfig()
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
