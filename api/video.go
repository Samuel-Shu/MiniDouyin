package api

import (
	"MiniDouyin/model"
	"MiniDouyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Feed 视频feed流
func Feed(c *gin.Context) {
	demoUser := model.User{
		Id:            1,
		Username:      "shuxin",
		FollowCount:   12,
		FollowerCount: 13,
		IsFollow:      false,
	}
	demo := []model.Video{{
		Id:            1,
		Author:        demoUser,
		PlayUrl:       "http://rp8cwyjwy.hn-bkt.clouddn.com/douyin.mp4?e=1675504751&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:A5FTVckQlut9mvix4y1tzmTUXMU=",
		CoverUrl:      "http://rp8cwyjwy.hn-bkt.clouddn.com/douyin.mp4?e=1675504751&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:A5FTVckQlut9mvix4y1tzmTUXMU=",
		FavoriteCount: 8555,
		CommentCount:  131311,
		IsFavorite:    false,
		Title:         "hello golang",
	},
	}
	video := model.VideoLists{
		Response: model.Response{
			StatusCode: utils.SUCCESS,
		},
		VideoList: demo,
		NextTime:  int32(time.Now().UnixNano()),
	}
	c.JSON(http.StatusOK, video)
}

// VideoPublish 视频发布
func VideoPublish(c *gin.Context)  {
	title := c.PostForm("title")
	token := c.PostForm("token")
	videoData,err := c.FormFile("data")
	utils.ResolveError(err)
	userId := utils.ParseToken(token)
	data := model.ParseVideo(videoData)
	key := fmt.Sprintf("video/%s.mp4",title)
	code := utils.PushVideo(key,data)
	if code == utils.SUCCESS{
		model.PushVideoToMysql(userId,utils.GetVideo(fmt.Sprintf("%s.mp4",title)),"",title)
		c.JSON(http.StatusOK,model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg: utils.GetStatusMsg(utils.VIDEO_PUSH_SUCCESS),
		})
	}else {
		c.JSON(http.StatusOK,model.Response{
			StatusCode: utils.FAIL,
			StatusMsg: utils.GetStatusMsg(utils.VIDEO_PUSH_FAIL),
		})
	}

}
