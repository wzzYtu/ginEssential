package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"ytu/ginessential/common"
	"ytu/ginessential/model"
)

func Login(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		c.JSON(422, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		c.JSON(422, gin.H{"code": 422, "msg": "密码必须大于6位"})
		return
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		c.JSON(422, gin.H{"code": 422, "msg": "用户不存在，请注册用户"})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	// 发送token
	token, err := common.ReleaseTocken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Print("token generate err:", err)
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}
