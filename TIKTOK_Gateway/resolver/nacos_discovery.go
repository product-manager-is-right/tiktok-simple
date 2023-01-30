package resolver

import (
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

func CreateDiscoveryClient() *client.Client {
	cli, err := client.NewClient()
	if err != nil {
		log.Fatal("创建服务发现客户端失败 : ", err.Error())
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("101.42.50.112", 8845),
	}
	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "info",
	}
	nacosCli, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
	r := nacos.NewNacosResolver(nacosCli)
	cli.Use(sd.Discovery(r))

	return cli
}
