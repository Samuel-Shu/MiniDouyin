package api

import (
	"MiniDouyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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
		CoverUrl:      "http://rp8cwyjwy.hn-bkt.clouddn.com/douyinJpg.jpg?e=1675506098&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:CScb-mrhpHGaB0RDiB-03RXa2PA=",
		FavoriteCount: 8555,
		CommentCount:  131311,
		IsFavorite:    false,
		Title:         "hello golang",
	},
	}
	video := model.VideoLists{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: demo,
		NextTime:  int32(time.Now().UnixNano()),
	}
	c.JSON(http.StatusOK, video)
}
