package model

import (
	"MiniDouyin/db"
	"testing"
)

func TestPushVideoToMysql(t *testing.T) {
	db.InitDb()
	PushVideoToMysql(1,"https://www.baidu.com","https://www.baidu.com","shu")
}
