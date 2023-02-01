package test

import (
	"GoProject/dal/mysql"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"log"
	"testing"
)

func TestA(t *testing.T) {
	mysql.Init()
	res, err := mysql.GetPublishVideoIdsById(24)
	if err != nil {
		log.Print(err)
		log.Print("find object failed")
	}
	fmt.Println(res)
	assert.Nil(t, nil)
}
