package model

import (
	"MiniDouyin/db"
	"testing"
)

func TestVideoIfFavorite(t *testing.T) {
	db.InitDb()
	db.Db.Model(&video{}).Where("video_id = ?",3).Update("favorite_count",0)
}
