package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"ytu/ginessential/common"
	"ytu/ginessential/model"
)

func UserRegister(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}
	if len(name) < 3 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "名称不能少于3位",
		})
		return
	}
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) == false {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已存在",
		})
		return
	}
	// 创建用户
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{Name: name, Telephone: telephone, Password: string(hasePassword)}
	DB.Create(&newUser)
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func isTelephoneExist(DB *gorm.DB, telephone string) bool {
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return false
	}
	return true
}
