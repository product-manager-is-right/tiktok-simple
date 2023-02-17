package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"TIKTOK_User/mw/rabbitMQ/producer"
	"TIKTOK_User/mw/redis"
	"TIKTOK_User/resolver"
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type FavoriteServiceImpl struct {
}

// CreateNewFavorite
/*
创建了一个favor对应关系
*/
func (fsi *FavoriteServiceImpl) CreateNewFavorite(userId, videoId int64) error {
	_, err := mysql.GetFavorite(userId, videoId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	// 数据库没有这条记录，插入
	if err == gorm.ErrRecordNotFound {
		err = producer.SendFavoriteMessage(userId, videoId, 1)
		if err != nil {
			log.Print("发送点赞操作消息队列失败，使用Mysql直接处理数据")
			_, err = mysql.CreateNewFavorite(userId, videoId)
			if err != nil {
				return err
			}
			// favorite数据库已经改变，删除redis userid对应的喜爱列表， 重试机制保证删除
			strUserId := strconv.FormatInt(userId, 10)
			for i := 0; i < redis.RetryTime; i++ {
				if _, err := redis.FavoriteList.Del(context.Background(), strUserId).Result(); err == nil {
					break
				}
			}

			go remoteUpdateFavoriteCnt(videoId, 0)
			return nil
		}
		return nil
	}
	log.Println("已有点赞无法再点赞")
	err = errors.New("已有点赞")
	return err
}

func (fsi *FavoriteServiceImpl) DeleteFavorite(userId, videoId int64) error {
	_, err := mysql.GetFavorite(userId, videoId)
	if err == gorm.ErrRecordNotFound {
		return errors.New("没有点赞过该视频，无法取消")
	}
	//先尝试发送到消息队列中
	if err = producer.SendFavoriteMessage(userId, videoId, 0); err != nil {
		if err = mysql.DeleteFavorite(userId, videoId); err != nil {
			return err
		}
		// favorite数据库已经改变，删除redis userid对应的喜爱列表， 重试机制保证删除
		strUserId := strconv.FormatInt(userId, 10)
		for i := 0; i < redis.RetryTime; i++ {
			if _, err := redis.FavoriteList.Del(context.Background(), strUserId).Result(); err == nil {
				break
			}
		}
		go remoteUpdateFavoriteCnt(videoId, 1)
	}

	return nil
}

// 远程调用video模块，修改video的点赞数
// actionType : 0->点赞  1->取消点赞
func remoteUpdateFavoriteCnt(videoId int64, actionType int) {
	client := resolver.GetNacosDiscoveryCli()
	args := &protocol.Args{}
	args.Add("video_id", strconv.Itoa(int(videoId)))
	if actionType == 0 {
		args.Add("action_type", "0")
	} else {
		args.Add("action_type", "1")
	}
	_, _, err := client.Post(context.Background(), nil, "http://tiktok.simple.video/douyin/video/favoriteAction/", args, config.WithSD(true))
	if err != nil {
		log.Println("远程调用点赞失败:" + err.Error())
	}
}

func (fsi *FavoriteServiceImpl) GetFavoriteVideosListByUserId(queryId, ownerId int64) ([]vo.VideoInfo, error) {
	var videoInfos []vo.VideoInfo

	// 先从redis中查找
	strQueryId := strconv.FormatInt(queryId, 10)
	vs, err := redis.FavoriteList.Get(context.Background(), strQueryId).Result()
	// 缓存命中
	if err == nil {
		// 反序列化
		if err := json.Unmarshal([]byte(vs), &videoInfos); err != nil {
			return nil, errors.New("redis Unmarshal:" + err.Error())
		}
		return videoInfos, nil
	}

	// 缓存未命中，查询数据库
	videoIds, err := mysql.GetFavoritesById(queryId)
	if err != nil {
		return nil, errors.New("数据库查询失败")
	}

	videoInfos, err = getVideoInfosByVideoIds(videoIds, ownerId, "favorite_query")
	if err != nil {
		return videoInfos, err
	}

	// 存入redis，不需要处理异常
	strVideoInfos, _ := json.Marshal(videoInfos)
	redis.FavoriteList.Set(context.Background(), strQueryId, strVideoInfos, redis.SetExpiredTime())

	return videoInfos, nil
}

// IsFavorite 断是否为喜欢接口,假如userId点赞了videoId，返回true,没有返回false
func (fsi *FavoriteServiceImpl) IsFavorite(userId, videoId int64) (bool, error) {
	isFavorite, err := mysql.GetIsFavorite(userId, videoId)
	if err != nil {
		log.Print("获取视频信息列表失败:", err.Error())
		return false, err
	}
	return isFavorite, nil
}
