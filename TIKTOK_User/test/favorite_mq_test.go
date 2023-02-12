package test

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/mw/rabbitMQ"
	"TIKTOK_User/service/serviceImpl"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
	"time"
)

func TestFavoriteMq(t *testing.T) {
	mysql.Init()
	rabbitMQ.InitRabbitMQ()
	rabbitMQ.InitFavoriteRabbitMQ()
	var err error
	var tests = []struct {
		userId  int64
		videoId int64
		cancel  bool
	}{
		{25, 9, true},
		{25, 9, false},
		//{25, 9, false},
	}
	fsi := serviceImpl.FavoriteServiceImpl{}
	for _, test := range tests {
		if test.cancel {
			if err = fsi.DeleteFavorite(test.userId, test.videoId); err != nil {
				t.Fatal(err)
			}
		} else {
			if err = fsi.CreateNewFavorite(test.userId, test.videoId); err != nil {
				t.Fatal(err)
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}

	time.Sleep(time.Second * 5)
	assert.Nil(t, err)
}
