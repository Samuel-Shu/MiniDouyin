package api

import (
	"MiniDouyin/middleware"
	"MiniDouyin/model"
	"MiniDouyin/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//GiveALikeWithVideo 视频点赞接口
func GiveALikeWithVideo(c *gin.Context)  {
	videoId := c.Query("video_id")
	videoIdInt, err2 := strconv.Atoi(videoId)
	if err2 !=nil{
		log.Fatal(err2)
	}
	actionType := c.Query("action_type")
	actionTypeInt,err := strconv.Atoi(actionType)
	if err!=nil{
		log.Fatal(err)
	}
	token := c.Query("token")
	tokenStr,err := middleware.ParseToken(token)
	if err != nil{
		log.Fatal(err)
	}

	//点赞功能
	if actionTypeInt == 1 {
		model.GiveALikeWithVideo(tokenStr.(middleware.MyClaim).UserName,tokenStr.(middleware.MyClaim).UserId,int32(videoIdInt))
	}else {
		//取消点赞
		model.UnlikeWithVideo(tokenStr.(middleware.MyClaim).UserId,int32(videoIdInt))
	}
	c.JSON(http.StatusOK,model.Response{
		StatusCode: utils.SUCCESS,
		StatusMsg: utils.GetStatusMsg(utils.VIDEO_GIVE_A_LIKE_SUCCESS),
	})
}
