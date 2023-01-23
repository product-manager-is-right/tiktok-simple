package model

type Video struct {
	Id            int64
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	PublishTime   int64
	Title         string
}

func (v *Video) TableName() string {
	return "vms_video"
}
