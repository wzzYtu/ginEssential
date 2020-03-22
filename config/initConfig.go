package config

import (
	"github.com/spf13/viper"
	"os"
)

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
