package main

import (
	"TIKTOK_Video/api/router"
	"TIKTOK_Video/configs"
	"TIKTOK_Video/dal"
	"TIKTOK_Video/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

func main() {
	// 读取配置文件
	configs.ReadConfig(configs.DEV)

	// 初始化工具
	initDeps()

	// 启动服务
	startServer()
}

func initDeps() {
	// 初始化数据库
	dal.Init()

	// 初始化Jwt
	mw.InitJwt()
}

func startServer() {
	cc := &constant.ClientConfig{
		Username: viper.GetString("nacos.username"),
		Password: viper.GetString("nacos.password"),
	}
	sc := []constant.ServerConfig{{
		IpAddr: viper.GetString("nacos.addr"),
		Port:   viper.GetUint64("nacos.port"),
	},
	}
	cli, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)

	r := nacos.NewNacosRegistry(cli)

	addr := ":" + viper.GetString("port")
	h := server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: viper.GetString("nacos.serviceName"),
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}))

	// 注册路由
	router.GeneratedRegister(h)

	h.Spin()
}
