package configs

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Port   int      `yaml:"port"`
	Nacos  *Nacos   `yaml:"nacos"`
	Routes []*Route `yaml:"routes"`
}
type Route struct {
	ServiceName string   `yaml:"serviceName"`
	Method      string   `yaml:"method"`
	Apis        []string `yaml:"apis"`
}

type Nacos struct {
	Addr string `yaml:"addr"`
	Port uint64 `yaml:"port"`
}

const (
	DEV  = "dev"
	TEST = "test"
)

func ReadConfig(mode string) (*Config, error) {
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
	// read the config
	in, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(in, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
