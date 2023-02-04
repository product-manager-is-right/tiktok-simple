package ServiceImpl

import (
	"TIKTOK_Video/model/vo"
	"TIKTOK_Video/resolver"
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/config"
	"log"
	"strconv"
)

type FavoriteServiceImpl struct {
}

var (
	IsFavoriteErr = errors.New("query isFavorite error")
)

// IsFavorite 判断是否为喜欢接口,假如userId点赞了videoId，返回true,没有返回false
func (fsi *FavoriteServiceImpl) IsFavorite(userId, videoId int64) (bool, error) {
	//创建请求
	// 判断是否为喜欢接口
	url := "http://tiktok.simple.user//douyin/favorite/IsFavor/?userId=" + strconv.FormatInt(userId, 10) + "&videoId=" + strconv.FormatInt(videoId, 10)
	client := resolver.GetNacosDiscoveryCli()
	state, body, err := client.Get(context.Background(), nil, url, config.WithSD(true))
	if err != nil || state != 200 {
		log.Fatal("请求User模块失败", err.Error())
		return false, IsFavoriteErr
	}
	var result = vo.FavoriteInfoResponse{}
	// TODO: 判断响应结果
	if err = json.Unmarshal(body, &result); err != nil {
		log.Fatal("请求User模块失败:", err.Error())
		return false, IsFavoriteErr
	}
	if result.StatusCode != 0 {
		log.Fatal("请求User模块失败:返回非0状态码")
		return false, IsFavoriteErr
	}
	return result.IsFavorite, nil
}
