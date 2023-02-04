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
	var path = ""
	if mode == DEV {
		path = "./configs/config.yaml"
		viper.SetConfigFile(path)
	} else if mode == TEST {
		path = "../configs/config.yaml"
		viper.SetConfigFile(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("配置文件未找到")
		} else {
			log.Fatal("配置文件找到了，但有其他错误")
		}
	}
}
