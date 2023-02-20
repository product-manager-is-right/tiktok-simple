package ServiceImpl

import (
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/model"
	"TIKTOK_Video/model/vo"
	"TIKTOK_Video/mw"
	"TIKTOK_Video/resolver"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	uuid "github.com/satori/go.uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

type VideoServiceImpl struct {
}

// GetVideoInfosByLatestTime /*
/*
根据最后传入的时间获取对应的视频
*/
func (vsi *VideoServiceImpl) GetVideoInfosByLatestTime(latestTime int64, userId int64) ([]vo.VideoInfo, int64, error) {
	var videoInfos []vo.VideoInfo
	nextTime := time.Now().UnixMilli()

	videos, err := mysql.GetVideosByTime(latestTime)
	if err != nil {
		return videoInfos, nextTime, err
	}

	if len(videos) == 0 {
		return videoInfos, nextTime, errors.New("video is empty")
	}

	videoInfos, err = bindVideoInfo(videos, userId)
	nextTime = videos[len(videos)-1].PublishTime

	return videoInfos, nextTime, err
}

/*
根据对应的video和userid返回video视频的list
*/
func bindVideoInfo(videos []*model.Video, userId int64) ([]vo.VideoInfo, error) {
	videoInfos := make([]vo.VideoInfo, len(videos))
	Ids := make([]int64, len(videos))

	for i, video := range videos {
		Ids[i] = video.UserId
	}
	// 远程调用，获取user/author的个人信息
	serviceImpl := UserServiceImpl{}
	authors, err := serviceImpl.GetUsersInfoByIds(Ids, userId)
	if err != nil {
		return nil, err
	}

	// 调用远程接口，判断userId是否喜欢videoId视频,但是调用时就和调用本地service一样，实现方式不一样
	// 需要与user通信，应定义到service层
	fsi := FavoriteServiceImpl{}
	var isFavorite bool
	for i, video := range videos {
		videoId := video.VideoId

		favoriteCount, _ := mysql.GetFavoriteCountByID(videoId)

		commentCount, _ := mysql.GetCommentCountByID(videoId)

		if userId >= 0 {
			isFavorite, _ = fsi.IsFavorite(userId, videoId)
		} else {
			isFavorite = false
		}

		videoInfos[i] = vo.VideoInfo{
			Id:            videoId,
			Author:        *authors[Ids[i]],
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
	}

	return videoInfos, nil
}

// PublishVideo /*
// 将视频存储到minio文件服务器
func (vsi *VideoServiceImpl) PublishVideo(userId int64, fileHeader *multipart.FileHeader, videoTitle string) error {
	// 处理multipart.FileHeader文件为byte[]
	file, err := fileHeader.Open()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}

	videoId, err := createVideoCall(userId, buf.Bytes(), videoTitle)
	if err != nil {
		return err
	}
	// 调用Dao层，存入ums_publish_video
	go remoteCreatePublishVideo(userId, videoId)

	return nil
}

// GetVideoInfosByIds /*
// 根据视频id获取videoInfo
func (vsi *VideoServiceImpl) GetVideoInfosByIds(Ids []int64) ([]vo.VideoInfo, error) {
	videos := make([]*model.Video, 0, len(Ids))
	for _, id := range Ids {
		v, err := mysql.GetVideoByID(id)
		if err != nil {
			continue
		}
		videos = append(videos, v)
	}
	VideoInfos, err := TransVideoInfo(videos)
	if err != nil {
		return nil, err
	}
	return VideoInfos, nil

}

// TransVideoInfo /*
// 将视频转换成Users模块需要的信息
func TransVideoInfo(videos []*model.Video) ([]vo.VideoInfo, error) {
	videoInfos := make([]vo.VideoInfo, len(videos))
	Ids := make([]int64, len(videos))

	for i, video := range videos {
		Ids[i] = video.UserId
	}
	// 远程调用，获取user/author的个人信息

	for i, video := range videos {
		videoId := video.VideoId

		favoriteCount, _ := mysql.GetFavoriteCountByID(videoId)

		commentCount, _ := mysql.GetCommentCountByID(videoId)

		// 需要与user通信，应定义到service层
		var favorite bool

		videoInfos[i] = vo.VideoInfo{
			Id:            videoId,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    favorite,
			Title:         video.Title,
		}
	}

	return videoInfos, nil
}

/*
remoteVideoCall
发起远程调用视频模块，存储video，返回videoId
@ return videoId
*/
func createVideoCall(userId int64, videoData []byte, videoTitle string) (int64, error) {
	//使用minio作为文件服务器
	fileReader := bytes.NewReader(videoData)
	// 随机生成文件名
	uu1 := uuid.NewV4().String()
	fileName := uu1 + "." + "mp4"
	uu2 := uuid.NewV4().String()
	pictureName := uu2 + "." + "jpg"
	// 两个不同的文件夹存储
	buketNameVideo := "tiktok"
	buketNamePicture := "tikpic"
	err := mw.UploadFile(buketNameVideo, fileName, fileReader, int64(len(videoData)), "video/mp4")
	if err != nil {
		log.Print("update File failed")
	}
	// 生成对应的视频url
	playUrl := "http://150.158.135.49:9000/" + buketNameVideo + "/" + fileName
	//使用ffmpeg切帧
	coverData, err := readFrameAsJpeg(playUrl)
	pictureReader := bytes.NewReader(coverData)
	//上传对应的picture到对应的文件夹
	if mw.UploadFile(buketNamePicture, pictureName, pictureReader, int64(len(coverData)), "image/jpeg") != nil {
		log.Print("update picture failed")
	}
	//生成对应的图片url
	coverUrl := "http://150.158.135.49:9000/" + buketNamePicture + "/" + pictureName
	//写入vm数据库
	videoId, err := mysql.CreateVideo(userId, playUrl, coverUrl, videoTitle)
	if err != nil {
		return 0, err
	}
	return videoId, nil
}

/*
视频切帧
*/
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	//根据对应的url切帧
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	//将img转化为buf用于图像上传
	if jpeg.Encode(buf, img, nil) != nil {
		return nil, err
	}
	//返回图像的byte[]流
	return buf.Bytes(), err
}

/*
远程调用User模块，将发布关系存入ums_publish_video表
*/
func remoteCreatePublishVideo(UserId, VideoId int64) {
	url := "http://tiktok.simple.user/douyin/publish/UserVideo/?userId=" + strconv.FormatInt(UserId, 10) + "&videoId=" + strconv.FormatInt(VideoId, 10)
	client := resolver.GetNacosDiscoveryCli()
	state, _, err := client.Post(context.Background(), nil, url, nil, config.WithSD(true))
	// check if the result is successful
	if state != consts.StatusOK || err != nil {
		log.Println("远程调用失败:" + err.Error())
	}
}
