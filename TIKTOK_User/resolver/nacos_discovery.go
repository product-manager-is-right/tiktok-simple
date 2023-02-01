package resolver

import (
	"TIKTOK_User/configs"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/hertz-contrib/registry/nacos"
)

var singleCli *client.Client

func GetInstance() *client.Client {
	if singleCli != nil {
		return singleCli
	}
	log.Println("服务发现失败，服务为空")
	return nil
}
func CreateDiscoveryServer() {
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	myConfig, err := configs.ReadConfig(configs.DEV)
	if err != nil {
		log.Fatal("文件读取"+
			"失败:", err.Error())
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(myConfig.Nacos.Addr, myConfig.Nacos.Port),
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
