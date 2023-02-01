package test

import (
	"TIKTOK_Video/configs"
	vo2 "TIKTOK_Video/model/vo"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"strconv"
	"testing"
)

func TestNacos(t *testing.T) {
	configs.ReadConfig(configs.TEST)
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
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
	if err != nil {
		panic(err)
	}
	r := nacos.NewNacosResolver(nacosCli)
	cli.Use(sd.Discovery(r))

	//status, body, err := cli.Post(context.Background(), nil, "http://tiktok.simple.user/douyin/user/login/?username=zhihao3&password=111111", nil, config.WithSD(true))
	args := &protocol.Args{}
	ids := make([]int, 0)
	ids = append(ids, 1)
	ids = append(ids, 28)
	ids = append(ids, 29)
	ids = append(ids, 30)
	ids = append(ids, 30)
	bytes, err := json.Marshal(ids)
	args.Add("user_ids", string(bytes))
	args.Add("user_id", strconv.Itoa(23))
	status, body, err := cli.Post(context.Background(), nil, "http://tiktok.simple.user/douyin/user/get", args, config.WithSD(true))
	res := vo2.Response{}
	if err = json.Unmarshal(body, &res); err != nil {
		return
	}
	users := make([]*vo2.UserInfo, len(ids))

	if err = json.Unmarshal([]byte(res.StatusMsg), &users); err != nil {
		return
	}
	fmt.Printf("%#v\n", users)
	fmt.Println(status)
	fmt.Println(string(body))
	assert.Nil(t, err)
}

//func TestUser(t *testing.T) {
//
//	if err != nil {
//		return
//	}
//	fmt.Println(string(bytes))
//	unIds := make([]int, 0)
//	bytes = []byte{'[', '1', '2', ',', '2', ',', '3', ',', '4', ']'}
//	err = json.Unmarshal(bytes, &unIds)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(unIds)
//
//}
