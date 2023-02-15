package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"TIKTOK_User/mw/rabbitMQ"
	"TIKTOK_User/resolver"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
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
		err = sendFavoriteMessage(userId, videoId, 1)
		if err != nil {
			log.Print("发送点赞操作消息队列失败，使用Mysql直接处理数据")
			_, err = mysql.CreateNewFavorite(userId, videoId)
			if err != nil {
				return err
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
	if err = sendFavoriteMessage(userId, videoId, 0); err != nil {
		if err = mysql.DeleteFavorite(userId, videoId); err != nil {
			return err
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

func (fsi *FavoriteServiceImpl) GetFavoriteVideosListByUserId(userIdTar, userIdSrc int64) ([]vo.VideoInfo, error) {
	var videoInfos []vo.VideoInfo
	videoIds, err := mysql.GetFavoritesById(userIdTar)
	if err != nil {
		log.Print("userIdTar=", userIdTar)
		//fmt.Print("获取视频id列表失败")
		return nil, err
	}

	if len(videoIds) == 0 {
		return nil, errors.New("获取视频id列表为空")
	}
	videoInfos, err = getVideoInfosByVideoIds(videoIds, userIdSrc, "favorite_query")
	if err != nil {
		log.Print("获取视频信息列表失败")
	}
	return videoInfos, err

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

// actionType为0时为将已有的数据库记录改为0，1改为1，2为创建新的数据行
// 发送结构为 userId-videoId-type的消息到rabbit中
func sendFavoriteMessage(userId int64, videoId int64, actionType int) error {
	//using rabbitMQ to store the info
	sb := strings.Builder{}
	//使用最高的36，压缩一下
	sb.WriteString(strconv.FormatInt(userId, 36))
	sb.WriteString("-")
	sb.WriteString(strconv.FormatInt(videoId, 36))
	sb.WriteString("-")
	sb.WriteString(strconv.Itoa(actionType))
	if err := rabbitMQ.RmqFavorite.PublishWithEx(sb.String()); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
