package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"ytu/ginessential/common"
	"ytu/ginessential/model"
	"ytu/ginessential/response"
)

func UserRegister(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	var user = model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		panic("BingJSON err")
	}
	name := user.Name
	telephone := user.Telephone
	password := user.Password
	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	if len(name) < 3 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "名称不能少于3位")
		return
	}
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) == false {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
	// 创建用户
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{Name: name, Telephone: telephone, Password: string(hasePassword)}
	DB.Create(&newUser)
	//返回结果
	response.Success(c, nil, "注册成功")
}

func isTelephoneExist(DB *gorm.DB, telephone string) bool {
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return false
	}
	return true
}
