package test

import (
	"GoProject/dal/mysql"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	mysql.Init()
	code := m.Run()

	os.Exit(code)
}
func TestA(t *testing.T) {
	res, err := mysql.GetPublishVideoIdsById(24)
	if err != nil {
		log.Print(err)
		log.Print("find object failed")
	}
	fmt.Println(res)
	assert.Nil(t, nil)
}
func TestFollow(t *testing.T) {
	res, err := mysql.GetFollowCntByUserId(25)
	fmt.Printf(strconv.FormatInt(res, 10))
	if err != nil {
		log.Print(err)
	}
	assert.Nil(t, nil)
}
func TestFollower(t *testing.T) {
	res, err := mysql.GetFollowerCntByUserId(24)
	fmt.Printf(strconv.FormatInt(res, 10))
	if err != nil {
		log.Print(err)
	}
	assert.Nil(t, nil)
}
func TestIsFollow(t *testing.T) {
	res, err := mysql.GetIsFollow(24, 25)
	if err != nil {
		log.Print(err)
	}
	if res {
		fmt.Println("successful")
	}
	assert.Nil(t, nil)
}
