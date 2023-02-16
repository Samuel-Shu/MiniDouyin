package model

import (
	"MiniDouyin/utils"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
)

type VideoLists struct {
	Response
	NextTime  int32   `json:"next_time"`
	VideoList []Video `json:"video_list"`
}

// ParseVideo 将*multipart.FileHeader类型转化为 []byte
func ParseVideo(videoData *multipart.FileHeader) []byte {
	file ,err := videoData.Open()
	utils.ResolveError(err)
	defer file.Close()
	data,err :=ioutil.ReadAll(file)
	if err != nil && err ==io.EOF{
		fmt.Println(err)
	}
	return data
}
