package configs

import (
	"gopkg.in/yaml.v3"
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

func ReadConfig() (*Config, error) {
	in, err := os.ReadFile("./configs/config.yaml")
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
