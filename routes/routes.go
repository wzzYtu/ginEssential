package routes

import (
	"github.com/gin-gonic/gin"
	"ytu/ginessential/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.UserRegister)
	r.POST("/api/auth/login", controller.Login)
	return r
}
