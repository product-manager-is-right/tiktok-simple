package test

import (
	"TIKTOK_Video/dal/mysql"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"os"
	"testing"
)

func TestGetCommentByVideoIds(t *testing.T) {

	comments, err := mysql.GetCommentByVideoIds(1)
	for _, pComment := range comments {
		fmt.Printf("%#v\n", *pComment)
	}
	assert.Nil(t, err)

}

func TestGetCommentByCommentId(t *testing.T) {
	comment, err := mysql.GetCommentByID(1)
	assert.Nil(t, err)
	fmt.Printf("%#v", *comment)

}

func TestMain(m *testing.M) {
	mysql.Init()
	code := m.Run()

	os.Exit(code)
}
