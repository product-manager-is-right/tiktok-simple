package ServiceImpl

import (
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/model"
	"TIKTOK_Video/model/vo"
	"TIKTOK_Video/mw"
	"bytes"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"
)

type VideoServiceImpl struct {
}

func (vsi *VideoServiceImpl) GetVideoInfosByLatestTime(latestTime int64, userName string) ([]vo.VideoInfo, int64, error) {
	var videoInfos []vo.VideoInfo
	nextTime := time.Now().UnixMilli()

	videos, err := mysql.GetVideosByTime(latestTime)
	if err != nil {
		return videoInfos, nextTime, err
	}

	if len(videos) == 0 {
		return videoInfos, nextTime, errors.New("video is empty")
	}
	videoInfos = make([]vo.VideoInfo, len(videos))
	videoInfos, err = bindVideoInfo(videoInfos, videos, userName)

	nextTime = videos[len(videos)-1].PublishTime
	return videoInfos, nextTime, err
}

func bindVideoInfo(videoInfos []vo.VideoInfo, videos []model.Video, userName string) ([]vo.VideoInfo, error) {
	for i, video := range videos {
		var err error
		videoId := video.VideoId

		favoriteCount, err := mysql.GetFavoriteCountByID(videoId)
		if err != nil {
			return videoInfos, err
		}
		commentCount, err := mysql.GetCommentCountByID(videoId)
		if err != nil {
			return videoInfos, err
		}
		// 需要与user通信，应定义到service层
		var favorite bool
		if userName != "" {
			favorite, err = isFavorite(videoId, userName)
			if err != nil {
				return videoInfos, err
			}
		}

		videoInfos[i].Id = videoId
		videoInfos[i].Author = vo.DemoUser
		//videoInfos[i].Author.Id = video.AuthorId
		videoInfos[i].PlayUrl = video.PlayUrl
		videoInfos[i].CoverUrl = video.CoverUrl
		videoInfos[i].FavoriteCount = favoriteCount
		videoInfos[i].CommentCount = commentCount
		videoInfos[i].IsFavorite = favorite
		videoInfos[i].Title = video.Title

	}

	return videoInfos, nil
}

// 调用远程接口，判断userName是否喜欢videoId视频
func isFavorite(videoId int64, userName string) (bool, error) {
	// TODO : impl
	return false, nil
}

func (vsi *VideoServiceImpl) PublishVideo(userId int64, fileHeader *multipart.FileHeader, videoTitle string) error {
	// 处理multipart.FileHeader文件为byte[]
	file, err := fileHeader.Open()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	// 远程调用video接口，并存入vms_publish_video
	videoId, err := CreateVideoCall(userId, buf.Bytes(), videoTitle)
	if err != nil {
		return err
	}
	// 调用Dao层，存入ums_publish_video
	if err = remoteCreatePublishVideo(userId, videoId); err != nil {
		return err
	}
	return nil
}

/*
remoteVideoCall
发起远程调用视频模块，存储video，返回videoId
@ return videoId
*/
func CreateVideoCall(userId int64, videoData []byte, videoTitle string) (int64, error) {
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
远程调用Video模块，获取每个Video的具体信息
*/
func remoteCreatePublishVideo(UserId, VideoId int64) error {
	// TODO : impl
	return nil
}
