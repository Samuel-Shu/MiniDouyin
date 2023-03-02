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

type comment struct {
	Id         int32      `json:"id"`          //评论id
	User       model.User `json:"user"`        //用户信息
	Content    string     `json:"content"`     //评论内容
	CreateDate string     `json:"create_date"` //创建时间
}

type commentResponse struct {
	model.Response
	Comment comment `json:"comment" `
}

type commentList struct {
	model.Response
	CommentList []comment `json:"comment_list"`
}

// PublishComment   发布评论
func PublishComment(c *gin.Context) {
	videoId := c.Query("video_id")
	commentText := c.Query("comment_text")
	actionType := c.Query("action_type")
	videoID, err := strconv.Atoi(videoId)
	token, err1 := middleware.ParseToken(c.Query("token"))
	if err1 != nil {
		log.Fatal(err1)
	}
	if err != nil {
		log.Fatal(err)
	}
	actionTypeInt, err := strconv.Atoi(actionType)
	if err != nil {
		log.Fatal(err)
	}
	if actionTypeInt == 1 {
		commentRes := model.PublishCommentToMysql(token.(middleware.MyClaim).UserId, int32(videoID), commentText)
		user := model.GetUserData(commentRes.Id)
		commentResponseToApp := commentResponse{
			Response: model.Response{
				StatusCode: utils.SUCCESS,
				StatusMsg:  utils.GetStatusMsg(utils.COMMENT_PUSH_SUCCESS),
			},
			Comment: comment{
				Id:         commentRes.CommentId,
				User:       user,
				Content:    commentRes.Content,
				CreateDate: commentRes.CreateTime,
			},
		}
		c.JSON(http.StatusOK, commentResponseToApp)
	} else {
		deleteCommentId := c.Query("comment_id")
		deleteCommentIdInt, _ := strconv.Atoi(deleteCommentId)
		model.DeleteCommentWithCommentId(int32(deleteCommentIdInt), int32(videoID))
	}
}

//GetCommentList  获取评论列表
func GetCommentList(c *gin.Context) {
	videoId := c.Query("video_id")
	videoIdInt, _ := strconv.Atoi(videoId)
	commentRes, count := model.GetCommentList(int32(videoIdInt))
	commentResList := make([]comment, count)
	for i := 0; i < int(count); i++ {
		user := model.GetUserData(commentRes[i].Id)
		commentResList[i].Id=commentRes[i].CommentId
		commentResList[i].User=user
		commentResList[i].Content=commentRes[i].Content
		commentResList[i].CreateDate=commentRes[i].CreateTime
	}
	c.JSON(http.StatusOK,commentList{
		Response:model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg: utils.GetStatusMsg(utils.COMMENT_LIST_GET_SUCCESS),
		},
		CommentList: commentResList,
	})
}
