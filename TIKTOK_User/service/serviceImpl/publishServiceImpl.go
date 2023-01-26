package serviceImpl

import (
	"GoProject/dal/mysql"
	"GoProject/model/vo"
	"GoProject/mw"
	"bytes"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"os"
)

type PublishServiceImpl struct {
}

func (psi *PublishServiceImpl) PublishVideo(userId int64, fileHeader *multipart.FileHeader, videoTitle string) error {
	// 处理multipart.FileHeader文件为byte[]
	file, err := fileHeader.Open()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	// 远程调用video接口，并存入vms_publish_video
	videoId, err := remoteCreateVideoCall(userId, buf.Bytes(), videoTitle)
	if err != nil {
		return err
	}
	// 调用Dao层，存入ums_publish_video
	if err = mysql.CreatePublishVideo(userId, videoId); err != nil {
		return err
	}
	return nil
}

/*
remoteVideoCall
发起远程调用视频模块，存储video，返回videoId
@ return videoId
*/
func remoteCreateVideoCall(userId int64, videoData []byte, videoTitle string) (int64, error) {
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
	playUrl := "http://120.25.2.146:9000/" + buketNameVideo + "/" + fileName
	//使用ffmpeg切帧
	coverData, err := readFrameAsJpeg(playUrl)
	pictureReader := bytes.NewReader(coverData)
	//上传对应的picture到对应的文件夹
	if mw.UploadFile(buketNamePicture, pictureName, pictureReader, int64(len(coverData)), "image/jpeg") != nil {
		log.Print("update picture failed")
	}
	//生成对应的图片url
	coverUrl := "http://120.25.2.146:9000/" + buketNamePicture + "/" + pictureName
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
		return videoInfos, errors.New("VideoList is empty")
	}
	videoInfos, err = getVideoInfosByVideoIds(videoIds, userIdSrc, "publish_query")
	return videoInfos, err
}

/*
通过videoIds封装VideoInfo，发布视频列表和喜欢列表可用
@ params

	videoIds: 视频ids
	userId: 登录的userId，用于查询登录账号与<视频(是否喜欢)和用户(是否关注)>的关系
*/
func getVideoInfosByVideoIds(videoIds []int64, userId int64, mode string) ([]vo.VideoInfo, error) {
	videoInfos, err := remoteGetVideoInfoCall(videoIds)
	if err != nil {
		return videoInfos, err
	}

	// 处理每个VideoInfo中的User和is_favorite
	usi := UserServiceImpl{}

	// 发布视频模式，发布的用户信息一致，无需重复查询
	if mode == "publish_query" {
		user, _ := usi.GetUserInfoById(videoInfos[0].Author.Id, userId)
		for i, videoInfo := range videoInfos {
			videoInfos[i].Author = user

			if isFavorite, _ := mysql.GetIsFavorite(userId, videoInfo.Id); isFavorite {
				videoInfos[i].IsFavorite = true
			}
		}
	} else {
		for i, videoInfo := range videoInfos {
			user, _ := usi.GetUserInfoById(videoInfo.Author.Id, userId)
			videoInfos[i].Author = user

			if isFavorite, _ := mysql.GetIsFavorite(userId, videoInfo.Id); isFavorite {
				videoInfos[i].IsFavorite = true
			}
		}
	}

	return videoInfos, nil
}

/*
远程调用Video模块，获取每个Video的具体信息
*/
func remoteGetVideoInfoCall(videoIds []int64) ([]vo.VideoInfo, error) {
	// TODO : impl
	return []vo.VideoInfo{}, nil
}
