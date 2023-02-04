package resolver

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"log"
	"sync"

	"github.com/cloudwego/hertz/pkg/app/client"
)

var cli *client.Client
var once = sync.Once{}

func GetNacosDiscoveryCli() *client.Client {
	once.Do(func() {
		var err error
		cli, _ = client.NewClient()

		sc := []constant.ServerConfig{
			*constant.NewServerConfig(viper.GetString("nacos.addr"), viper.GetUint64("nacos.port")),
		}
		cc := constant.ClientConfig{
			NamespaceId:         "public",
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogLevel:            "info",
		}
		nacosCli, err := clients.NewNamingClient(
			vo.NacosClientParam{
				ClientConfig:  &cc,
				ServerConfigs: sc,
			})
		if err != nil {
			log.Println(err.Error())
		}
		r := nacos.NewNacosResolver(nacosCli)

		cli.Use(sd.Discovery(r))
	})
	return cli
}
