package main

import (
	"TIKTOK_User/api/router"
	"TIKTOK_User/dal"
	"TIKTOK_User/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	initDeps()

	cc := &constant.ClientConfig{
		AppName:  "test",
		Username: "nacos",
		Password: "nacos",
	}
	sc := []constant.ServerConfig{{
		IpAddr: "101.42.50.112",
		Port:   8848,
	},
	}
	cli, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)

	r := nacos.NewNacosRegistry(cli)

	h := server.Default(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "tiktok.simple.user",
			Addr:        utils.NewNetAddr("tcp", "127.0.0.1:8080"),
			Weight:      10,
			Tags:        nil,
		}))

	// 注册路由
	router.GeneratedRegister(h)

	h.Spin()
}

func initDeps() {
	// 初始化数据库
	dal.Init()

	// 初始化Jwt
	mw.InitJwt()
}
