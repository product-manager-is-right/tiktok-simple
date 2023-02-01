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
var once sync.Once = sync.Once{}

func GetInstance() *client.Client {
	if cli != nil {
		return cli
	}
	log.Println("服务发现失败，服务为空")
	return nil
}
func CreateDiscoveryServer() {
	once.Do(func() {
		var err error
		cli, err = client.NewClient()
		if err != nil {
			panic(err)
		}
		//myConfig, err := configs.ReadConfig(configs.DEV)
		if err != nil {
			log.Fatal("文件读取"+
				"失败:", err.Error())
		}
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
		r := nacos.NewNacosResolver(nacosCli)
		if err != nil {
			log.Fatal(err)
		}
		cli.Use(sd.Discovery(r))
	})

}

/*
	for i := 0; i < 10; i++ {
		status, body, err := cli.Get(context.Background(), nil, "http://hertz.test.demo/ping", config.WithSD(true))
		if err != nil {
			hlog.Fatal(err)
		}
		hlog.Infof("code=%d,body=%s\n", status, string(body))
	}
*/
