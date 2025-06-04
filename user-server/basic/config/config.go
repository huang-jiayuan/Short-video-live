package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Mysql struct {
		User     string
		Password string
		Host     string
		Port     int
		Database string
	}
	Redis struct {
		Host     string
		Password string
		Db       int
	}
}

var Appconf Config

func Inits() {
	viper.SetConfigFile("./basic./config./dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("文件读取成功")
	err = viper.Unmarshal(&Appconf)
	if err != nil {
		panic(err)
	}
	fmt.Println(Appconf)
}
