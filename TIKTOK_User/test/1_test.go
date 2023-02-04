package test

import (
	"TIKTOK_User/configs"
	"TIKTOK_User/dal/mysql"
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
	"log"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	configs.ReadConfig(configs.TEST)

	mysql.Init()
	code := m.Run()

	os.Exit(code)
}
func TestA(t *testing.T) {
	res, err := mysql.GetPublishVideoIdsById(24)
	if err != nil {
		log.Print(err)
		log.Print("find object failed")
	}
	fmt.Println(res)
	assert.Nil(t, nil)
}
func TestFollow(t *testing.T) {
	res, err := mysql.GetFollowCntByUserId(25)
	fmt.Printf(strconv.FormatInt(res, 10))
	if err != nil {
		log.Print(err)
	}
	assert.Nil(t, nil)
}
func TestFollower(t *testing.T) {
	res, err := mysql.GetFollowerCntByUserId(24)
	fmt.Printf(strconv.FormatInt(res, 10))
	if err != nil {
		log.Print(err)
	}
	assert.Nil(t, nil)
}
func TestIsFollow(t *testing.T) {
	res, err := mysql.GetIsFollow(24, 25)
	if err != nil {
		log.Print(err)
	}
	if res {
		fmt.Println("successful")
	}
	assert.Nil(t, nil)
}

func TestGetVideos(t *testing.T) {
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
	ids := make([]int64, 0)
	ids = append(ids, 11)
	ids = append(ids, 12)
	ids = append(ids, 13)
	bytes, err := json.Marshal(ids)
	args.Add("videoIds", string(bytes))
	status, body, err := cli.Post(context.Background(), nil, "http://tiktok.simple.video/douyin/publish/GetVideos/", args, config.WithSD(true))
	fmt.Println(status)
	fmt.Println(string(body))
	assert.Nil(t, err)
}
