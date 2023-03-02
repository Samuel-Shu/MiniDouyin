package model

import (
	"MiniDouyin/db"
	"gorm.io/gorm"
)

type Favorite struct {
	FavoriteId int32  `json:"favorite_id"`
	Username   string `json:"username"`
	Id         int32  `json:"id"`
	VideoId    int32  `json:"video_id"`
	IsDeleted  int32  `json:"is_deleted"`
}

//GiveALikeWithVideo 根据videoId以及userId实现视频点赞
func GiveALikeWithVideo(username string, id, videoId int32) {
	favorite := Favorite{
		Username: username,
		Id:       id,
		VideoId:  videoId,
	}
	if !VideoIfFavoriteInGiveALike(id, videoId) {
		db.Db.Create(&favorite)
		db.Db.Model(&video{}).Where("video_id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
	} else {
		db.Db.Model(&Favorite{}).Where("id = ? && video_id = ?", id, videoId).Update("is_deleted", gorm.Expr("is_deleted - ?", 1))
		db.Db.Model(&video{}).Where("video_id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
	}
}

//UnlikeWithVideo 根据videoId以及userId取消视频点赞
func UnlikeWithVideo(id, videoId int32) {
	db.Db.Model(&Favorite{}).Where("id = ? && video_id = ?", id, videoId).Update("is_deleted", 1)
	db.Db.Model(&video{}).Where("video_id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
}

//VideoIfFavoriteInGiveALike 判断用户是否已经点赞  true--已点赞；false--未点赞
func VideoIfFavoriteInGiveALike(id, videoId int32) bool {
	favorite := Favorite{}
	db.Db.Where("id = ? && video_id = ?", id, videoId).Find(&favorite)
	if favorite.FavoriteId != 0 {
		return true
	}
	return false
}

//VideoIfFavorite 判断用户是否已经点赞  true--已点赞；false--未点赞
func VideoIfFavorite(id, videoId int32) bool {
	favorite := Favorite{}
	db.Db.Where("id = ? && video_id = ?", id, videoId).Find(&favorite)
	if favorite.FavoriteId != 0 && favorite.IsDeleted == 0 {
		return true
	}
	return false
}

//GetFavoriteVideoList 根据user_id获取喜欢视频
func GetFavoriteVideoList(userId int32) ([]video, int64) {
	var favoriteList []Favorite
	var countFavorite int64
	var VideoList []video
	db.Db.Where("id = ? && is_deleted = ?", userId, 0).Find(&favoriteList).Count(&countFavorite)
	for i := 0; i < int(countFavorite); i++ {
		db.Db.Where("video_id = ?", favoriteList[i].VideoId).Find(&VideoList)
	}
	return VideoList, countFavorite
}
