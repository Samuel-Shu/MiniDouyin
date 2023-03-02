package api

import (
	"MiniDouyin/middleware"
	"MiniDouyin/model"
	"MiniDouyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// VideoPublish 视频发布
func VideoPublish(c *gin.Context) {
	title := c.PostForm("title")
	token := c.PostForm("token")
	videoData, err := c.FormFile("data")
	utils.ResolveError(err)

	claim, err := middleware.ParseToken(token)
	if err != nil {
		log.Fatal(err)
	}
	userId := claim.(middleware.MyClaim).UserId

	data := model.ParseVideo(videoData)

	keyVideo := fmt.Sprintf("%s.mp4", title)
	codeVideo := utils.PushVideo(keyVideo, data)
	playUrl := utils.GetVideo(fmt.Sprintf("%s.mp4", title))
	coverByte, err := utils.ParseCover(playUrl, 1)
	keyPicture := fmt.Sprintf("%s.jpg", title)
	utils.ResolveError(err)
	codePicture := utils.PushVideoCover(keyPicture, coverByte)
	coverUrl := utils.GetCover(keyPicture)
	if codeVideo == utils.SUCCESS && codePicture == utils.SUCCESS {
		model.PushVideoToMysql(userId, playUrl, coverUrl, title)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg:  utils.GetStatusMsg(utils.VIDEO_PUSH_SUCCESS),
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: utils.FAIL,
			StatusMsg:  utils.GetStatusMsg(utils.VIDEO_PUSH_FAIL),
		})
	}

}

//Feed 视频feed流
func Feed(c *gin.Context) {
	var timeUnix time.Time
	latestTime := c.Query("latest_time")

	//根据latestTime来查找视频feed流
	videoData, count := model.GetVideo(latestTime)
	videoAndAuthor := make([]model.Video, count)

	//将拿到的feed流遍历写入VideoAndAuthor
	for i := 0; i < int(count); i++ {
		user := model.GetUserData(videoData[i].Id)
		videoAndAuthor[i].Id = videoData[i].VideoId
		videoAndAuthor[i].PlayUrl = videoData[i].PlayUrl
		videoAndAuthor[i].CoverUrl = videoData[i].CoverUrl
		videoAndAuthor[i].CommentCount = videoData[i].CommentCount
		videoAndAuthor[i].FavoriteCount = videoData[i].FavoriteCount
		if model.VideoIfFavorite(videoData[i].Id, videoData[i].VideoId) {
			videoAndAuthor[i].IsFavorite = true
		} else {
			videoAndAuthor[i].IsFavorite = false
		}
		videoAndAuthor[i].Title = videoData[i].Title
		videoAndAuthor[i].CreateDate = videoData[i].CreateDate
		videoAndAuthor[i].Author = user
	}

	// 确定nextTime的时间戳，后续不存在视频则时间戳为当前时间，存在则为最早投稿的时间戳
	if count == 0 {
		timeUnix = time.Now()
	} else {
		timeUnix, _ = time.Parse("2006-01-02T15:04:05Z07:00", videoAndAuthor[count-1].CreateDate)
	}

	//feed流返回信息
	videoLists := model.VideoLists{
		Response: model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg:  utils.GetStatusMsg(utils.VIDEO_GET_SUCCESS),
		},
		NextTime:  timeUnix.Unix(),
		VideoList: videoAndAuthor,
	}

	c.JSON(http.StatusOK, videoLists)
}

// GetVideoList 获取用户视频发布列表
func GetVideoList(c *gin.Context) {
	UserId := c.Query("user_id")
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		log.Fatal(userId)
	}
	videoData, count := model.GetVideoList(int32(userId))
	videoAndAuthor := make([]model.Video, count)
	for i := 0; i < int(count); i++ {
		user := model.GetUserData(videoData[i].Id)
		videoAndAuthor[i].Id = videoData[i].VideoId
		videoAndAuthor[i].PlayUrl = videoData[i].PlayUrl
		videoAndAuthor[i].CoverUrl = videoData[i].CoverUrl
		videoAndAuthor[i].CommentCount = videoData[i].CommentCount
		videoAndAuthor[i].FavoriteCount = videoData[i].FavoriteCount
		videoAndAuthor[i].IsFavorite = videoData[i].IsFavorite
		videoAndAuthor[i].Title = videoData[i].Title
		videoAndAuthor[i].CreateDate = videoData[i].CreateDate
		videoAndAuthor[i].Author = user
	}
	videoLists := model.VideoLists{
		Response: model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg:  utils.GetStatusMsg(utils.VIDEO_GET_SUCCESS),
		},
		VideoList: videoAndAuthor,
	}
	c.JSON(http.StatusOK, videoLists)
}
