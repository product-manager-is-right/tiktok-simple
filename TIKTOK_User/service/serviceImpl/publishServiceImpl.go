package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"TIKTOK_User/resolver"
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	"log"
	"math/rand"
	"time"
)

type PublishServiceImpl struct {
}

// ErrGetVideosInfo
// 错误返回值
var ErrGetVideosInfo = errors.New("can not get the VideoInfo")

/*
GetVideoList
获取userId发布的视频列表

@ params

	userIdTar: 请求查询的userId
	userIdSrc: 登录userId
*/
func (psi *PublishServiceImpl) GetVideoList(userIdTar int64, userIdSrc int64) ([]vo.VideoInfo, error) {
	// 调用Dao层，查找userIdTar用户的所有videoIds
	var videoInfos []vo.VideoInfo
	videoIds, err := mysql.GetPublishVideoIdsById(userIdTar)
	if err != nil {
		return videoInfos, err
	}
	if len(videoIds) == 0 {
		return videoInfos, nil
	}

	videoInfos, err = getVideoInfosByVideoIds(videoIds, userIdSrc, "publish_mode")

	return videoInfos, err
}

/*
通过videoIds封装VideoInfo，发布视频列表和喜欢列表可用
@ params

	videoIds: 视频ids
	userId: 登录的userId，用于查询登录账号与<视频(是否喜欢)和用户(是否关注)>的关系
	mode: "publish_mode" or "favorite_mode" 发布视频列表查询模式/喜欢列表查询模式
*/
func getVideoInfosByVideoIds(videoIds []int64, userId int64, mode string) ([]vo.VideoInfo, error) {
	videoInfos, err := remoteGetVideoInfoCall(videoIds)
	if err != nil {
		return videoInfos, err
	}

	// 处理每个VideoInfo中的User和is_favorite
	usi := UserServiceImpl{}

	// 发布视频模式，发布的用户信息一致，无需重复查询
	if mode == "publish_mode" {
		user, _ := usi.GetUserInfoById(videoInfos[0].Author.Id, userId)
		for i, videoInfo := range videoInfos {
			videoInfos[i].Author = *user

			if isFavorite, _ := mysql.GetIsFavorite(userId, videoInfo.Id); isFavorite {
				videoInfos[i].IsFavorite = true
			}
		}
	} else if mode == "favorite_mode" {
		for i, videoInfo := range videoInfos {
			user, _ := usi.GetUserInfoById(videoInfo.Author.Id, userId)
			videoInfos[i].Author = *user

			// 喜欢列表模式，查询的视频列表都是已赞
			videoInfos[i].IsFavorite = true
		}
	}

	return videoInfos, nil
}

/*
远程调用Video模块，获取每个Video的具体信息
*/
func remoteGetVideoInfoCall(videoIds []int64) ([]vo.VideoInfo, error) {
	// TODO : remote impl
	vs := make([]vo.VideoInfo, len(videoIds))
	rand.Seed(time.Now().Unix())
	client := resolver.GetInstance()
	args := &protocol.Args{}
	bytes, err := json.Marshal(videoIds)
	if err != nil {
		log.Print("failed to change videoIds to json")
	}
	//用 bytes的方式存储videos的id
	args.Add("video", string(bytes))
	status, body, err := client.Post(context.Background(), nil, "http://tiktok.simple.video/douyin/publish/GetVideos", args, config.WithSD(true))
	if status == 200 {
		res := vo.Response{}
		if err = json.Unmarshal(body, &res); err != nil {
			return nil, ErrGetVideosInfo
		}
		videos := make([]*vo.VideoInfo, len(videoIds))
		if err = json.Unmarshal([]byte(res.StatusMsg), &videos); err != nil {
			return nil, ErrGetVideosInfo
		}
		for index, video := range videos {
			vs[index] = vo.VideoInfo{
				Id:            video.Id,
				PlayUrl:       video.PlayUrl,
				CoverUrl:      video.CoverUrl,
				FavoriteCount: video.FavoriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    video.IsFavorite,
				Title:         "Test",
			}

		}
	}
	return vs, nil
	/*
		for i, videoId := range videoIds {
				vs[i] = vo.VideoInfo{
					Id:            videoId,
					PlayUrl:       "http://120.25.2.146:9000/tiktok/videos/test.mp4",
					CoverUrl:      "http://120.25.2.146:9000/tiktok/picture/testP.jpg",
					FavoriteCount: rand.Int63() % 10000,
					CommentCount:  rand.Int63() % 10000,
					IsFavorite:    videoId%73 == 0,
					Title:         "Test",
				}
			}
	*/
}

// PublishVideoInfo /*
// 检查是否存在用户id并存储该videoId 和 UserId 到 ums数据库
func (psi *PublishServiceImpl) PublishVideoInfo(userId, videoId int64) error {
	//检查是否有该userid
	_, err := mysql.GetUserByUserId(userId)
	if err != nil {
		return err
	}
	//创建video和user进入数据库
	if err := mysql.CreatePublishVideo(userId, videoId); err != nil {
		return err
	}
	return nil
}
