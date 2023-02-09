package router

import (
	"MiniDouyin/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	basicAPI := r.Group("/douyin")
	basicAPI.POST("/user/register/", api.Register)
	basicAPI.GET("/user/", api.GetUserData)
	basicAPI.POST("/user/login/", api.Login)
	basicAPI.GET("/feed/", api.Feed)
	err := r.Run(":8088")
	if err != nil {
		panic(err)
	}
}
