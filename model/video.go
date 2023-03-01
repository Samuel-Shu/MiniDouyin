package model

import (
	"MiniDouyin/config"
	"MiniDouyin/db"
	"MiniDouyin/utils"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"strconv"
	"time"
)

type VideoLists struct {
	Response
	NextTime  int64   `json:"next_time"`
	VideoList []Video `json:"video_list"`
}

type video struct {
	VideoId       int32
	Id            int32
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int32
	CommentCount  int32
	Title         string
	IsFavorite    bool
	CreateDate    string
}

// ParseVideo 将*multipart.FileHeader类型转化为 []byte
func ParseVideo(videoData *multipart.FileHeader) []byte {
	file, err := videoData.Open()
	utils.ResolveError(err)
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil && err == io.EOF {
		fmt.Println(err)
	}
	return data
}

// PushVideoToMysql 将投稿信息存入数据库
func PushVideoToMysql(id int32, playUrl, coverUrl, title string) {
	pushVideoToMysql := video{
		Id:       id,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
		CreateDate: time.Now().UTC().Format("2006-01-02 15:04:05"),
	}
	db.Db.Create(pushVideoToMysql)
}

//GetVideo 按照time降序的方式查找config.N个视频信息
func GetVideo(latestTime string) ([]video, int32) {
	var count int64
	var video []video
	timeInt, err := strconv.ParseInt(latestTime, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	timeStr := time.Unix(timeInt,0).Format("2006-01-02 15:04:05")
	db.Db.Where("create_date < ?", timeStr).Limit(config.N).Find(&video).Count(&count)
	if count >= config.N {
		count = config.N
	}
	return video, int32(count)
}

//GetVideoList 根据user_id查询发布视频列表
func GetVideoList(userId int32) ([]video,int64) {
	var videoList []video
	var count int64
	db.Db.Where("id=?",userId).Find(&videoList).Count(&count)
	return videoList,count
}
