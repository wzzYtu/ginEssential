package routes

import (
	"github.com/gin-gonic/gin"
	"ytu/ginessential/controller"
	"ytu/ginessential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.POST("/api/auth/register", controller.UserRegister)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
