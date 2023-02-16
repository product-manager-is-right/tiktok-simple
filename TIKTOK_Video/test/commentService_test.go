package test

import (
	"TIKTOK_Video/service"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestGetCommentListByVideoId(t *testing.T) {
	instance := service.NewCommentServiceInstance()
	var tests = []struct {
		videoId   int64
		ownerId   int64
		expectNil bool
	}{
		{1, 1, true},
		{2, 0, true},
		{-1, 0, false},
	}
	for _, test := range tests {
		commentInfos, err := instance.GetCommentListByVideoId(test.videoId, test.ownerId)
		if test.expectNil {
			assert.Nil(t, err)
			fmt.Println()
			for _, commentInfo := range commentInfos {
				fmt.Printf("%#v\n", *commentInfo)
			}
			fmt.Println()
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestInsertComment(t *testing.T) {
	var tests = []struct {
		videoId   int64
		ownerId   int64
		comment   string
		expectNil bool
		errorMsg  string
	}{
		{1, 3, "user3 inserted by test", true, ""},
		{-1, 3, "user3 inserted by test", false, "zero row affected"},
	}
	instance := service.NewCommentServiceInstance()
	for _, test := range tests {
		commentInfo, err := instance.InsertComment(test.comment, test.videoId, test.ownerId)
		if test.expectNil {
			assert.Nil(t, err)
			//assert.DeepEqual(t,test.errorMsg,err.Error())
			fmt.Println()
			fmt.Printf("%#v\n", *commentInfo)
			fmt.Println()
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestDeleteCommentByCommentIdWithMock(t *testing.T) {
	var tests = []struct {
		ownerId   int64
		commentId int64
		viderId   int64
		expectNil bool
		errorMsg  string
	}{
		//{3, 12, true, ""},
		{3, 27, 3, true, ""},
		//{6, 6, false, "zero row affected"},
	}
	instance := service.NewCommentServiceInstance()
	// todo 本来想像这样子打桩的，但是好像都没有打桩成功,可能是函数内联的原因，听说使用gomock好点
	//monkey.Patch(instance.ReturnA, func() string {
	//	return "b"
	//})
	//defer monkey.Unpatch(instance.ReturnA)
	//fmt.Println(instance.ReturnA())
	//monkey.Patch(instance.DeleteCommentByCommentId, func(commentId, userId int64) error {
	//	fmt.Printf("mock")
	//	if commentId == 12 && userId == 3 {
	//		return nil
	//	}
	//	return errors.New("zero row affected")
	//})
	//defer monkey.Unpatch(instance.DeleteCommentByCommentId)
	for _, test := range tests {
		err := instance.DeleteCommentByCommentId(test.commentId, test.ownerId, test.viderId)
		if test.expectNil {
			assert.Nil(t, err)
			//assert.DeepEqual(t,test.errorMsg,err.Error())
		} else {
			assert.NotNil(t, err)
			fmt.Println(err)
		}
	}
}

//func TestMain(m *testing.M) {
//
//}
