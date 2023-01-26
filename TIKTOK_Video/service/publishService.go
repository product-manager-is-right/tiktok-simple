package service

type PublishService interface {
	PublishVideo(userId int64, videoData []byte, videoTitle string) error
}
