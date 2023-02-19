package model

import (
	"MiniDouyin/db"
	"fmt"
	"time"
)

func timeDemo()  {
	video := video{}
	db.InitDb()
	db.Db.Where("video_id = ?",1).First(&video)
	time1,_:=time.Parse("2006-01-02 15:04:05",video.CreateDate)
	timeUnix := time1.Unix()
	fmt.Println(timeUnix)
	timeStr := time.Unix(timeUnix,10).Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)
}
