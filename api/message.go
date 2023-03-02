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

type messageListInRes struct {
	model.Response
	MessageList []model.Message `json:"message_list"`
}


//PushMessage  发送信息
func PushMessage(c *gin.Context)  {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")
	actionType := c.Query("action_type")
	parseToken, err := middleware.ParseToken(token)
	if err != nil{
		log.Fatal(err)
	}
	toUserIdInt, err1 := strconv.Atoi(toUserId)
	if err1 != nil{
		log.Fatal(err1)
	}
	actionTypeInt, err2 := strconv.Atoi(actionType)
	if err2!=nil{
		log.Fatal(err2)
	}
	if actionTypeInt ==1{
		model.SendMessage(parseToken.(middleware.MyClaim).UserId, int32(toUserIdInt),content)
		c.JSON(http.StatusOK,model.Response{StatusCode: utils.SUCCESS,StatusMsg: utils.GetStatusMsg(utils.SUCCESS)})
	}else {
		c.JSON(http.StatusOK,model.Response{StatusCode: utils.FAIL,StatusMsg: utils.GetStatusMsg(utils.FAIL)})
	}
}

//GetMessageChat  获取聊天记录
func GetMessageChat(c *gin.Context)  {
	var messageRes messageListInRes
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	parseToken, err := middleware.ParseToken(token)
	if err !=nil{
		log.Fatal(err)
	}
	toUserIdInt, err := strconv.Atoi(toUserId)
	if err != nil{
		log.Fatal(err)
	}
	messageList := model.GetMessage(parseToken.(middleware.MyClaim).UserId,int32(toUserIdInt))
	messageRes = messageListInRes{
		Response:model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg: utils.GetStatusMsg(utils.SUCCESS),
		},
		MessageList: messageList,
	}
	c.JSON(http.StatusOK,messageRes)
}