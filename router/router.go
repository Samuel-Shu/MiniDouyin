package router

import (
	"MiniDouyin/api"
	"MiniDouyin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	/* @ShuCoding
	basic service：
		register、login、feed、getUserData
	 */
	basicAPI := r.Group("/douyin")
	basicAPI.POST("/user/register/", api.Register)
	basicAPI.GET("/user/",middleware.JWT(), api.GetUserData)
	basicAPI.POST("/user/login/", api.Login)
	basicAPI.GET("/feed/", api.Feed)
	basicAPI.POST("/publish/action/",middleware.JWT(),api.VideoPublish)
	basicAPI.GET("/publish/list/",middleware.JWT(),api.GetVideoList)
	/*
	interact service :
	 */
	interactAPI := r.Group("/douyin")
	interactAPI.POST("/favorite/action/",middleware.JWT(),api.GiveALikeWithVideo)
	interactAPI.GET("/favorite/list/",middleware.JWT(),api.GetFavoriteVideoList)
	interactAPI.POST("/comment/action/",middleware.JWT(),api.PublishComment)
	interactAPI.GET("/comment/list/",api.GetCommentList)

	/*
	social service:
	 */
	socialAPI := r.Group("/douyin")
	socialAPI.Use(middleware.JWT())
	socialAPI.POST("/relation/action/",api.AttentionAction)
	socialAPI.GET("/relation/follow/list/",api.GetAttentionList)
	socialAPI.GET("/relation/follower/list/",api.GetFollowerList)
	socialAPI.GET("/relation/friend/list/",api.GetFriendList)
	socialAPI.POST("/message/action/",api.PushMessage)
	socialAPI.GET("/message/chat/",api.GetMessageChat)


	err := r.Run(":8088")
	if err != nil {
		panic(err)
	}
}
