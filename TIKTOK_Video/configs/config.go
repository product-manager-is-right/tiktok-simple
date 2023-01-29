package configs

import (
	"github.com/spf13/viper"
	"log"
)

const (
	DEV  = "dev"
	TEST = "test"
)

func ReadConfig(mode string) {
	if mode == DEV {
		viper.SetConfigFile("./configs/config.yaml")
	} else if mode == TEST {
		viper.SetConfigFile("../configs/config.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("配置文件未找到")
		} else {
			println(err.Error())
			log.Fatal("配置文件找到了，但有其他错误")
		}
	}

	validation()
}

func validation() {
	return
}
