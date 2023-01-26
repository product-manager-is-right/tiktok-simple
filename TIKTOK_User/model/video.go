package model

type Video struct {
	Id            int64
	UserId        int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	PublishTime   int64
	Title         string
}
