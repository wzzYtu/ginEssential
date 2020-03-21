package common

import (
	"github.com/jinzhu/gorm"
	"ytu/ginessential/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/essential?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
