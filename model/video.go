package model

type VideoLists struct {
	Response
	NextTime  int32   `json:"next_time"`
	VideoList []Video `json:"video_list"`
}
