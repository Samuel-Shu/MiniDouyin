package model

import (
	"MiniDouyin/db"
	"fmt"
	"testing"
)

func TestPushVideoToMysql(t *testing.T) {
	db.InitDb()
	n, count := GetVideo("2023-02-18 20:05")
	fmt.Println(n)
	fmt.Println(count)
}
