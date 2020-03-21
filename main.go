package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"ytu/ginessential/common"
	"ytu/ginessential/routes"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = routes.CollectRoute(r)
	r.Run()
}
