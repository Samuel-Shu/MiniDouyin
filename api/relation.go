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

type attentionList struct {
	model.Response
	UserList []model.User `json:"user_list"`
}

//AttentionAction  关注用户操作
func AttentionAction(c *gin.Context)  {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	parseToken, err := middleware.ParseToken(token)
	if err != nil{
		log.Fatal(err)
	}
	userId := parseToken.(middleware.MyClaim).UserId
	toUserIdInt,_ := strconv.Atoi(toUserId)
	actionTypeInt ,_ := strconv.Atoi(actionType)
	if actionTypeInt == 1 {
		model.AttentionActionToMysql(userId, int32(toUserIdInt))
	}else {
		model.DeleteActionToMysql(userId, int32(toUserIdInt))
	}
	c.JSON(http.StatusOK,model.Response{
		StatusCode: utils.SUCCESS,
		StatusMsg: utils.GetStatusMsg(utils.ATTENTION_SUCCESS),
	})
}

//GetAttentionList  获取关注列表
func GetAttentionList(c *gin.Context)  {
	userId := c.Query("user_id")
	userIdInt ,_ := strconv.Atoi(userId)
	relation,count := model.GetRelationListWithFollower(int32(userIdInt))
	userList := make([]model.User,count)
	for i:=0;i<int(count);i++{
		user := model.GetUserData(relation[i].FollowingId)
		userList[i]=user
	}
	c.JSON(http.StatusOK,attentionList{
		Response:model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg: utils.GetStatusMsg(utils.GET_ATTENTION_LIST_SUCCESS),
		},
		UserList: userList,
	})
}

//GetFollowerList  获取粉丝列表
func GetFollowerList(c *gin.Context)  {
	userId := c.Query("user_id")
	userIdInt ,_ := strconv.Atoi(userId)
	relation,count := model.GetRelationListWithFollowing(int32(userIdInt))
	userList := make([]model.User,count)
	for i:=0;i<int(count);i++{
		user := model.GetUserData(relation[i].FollowingId)
		userList[i]=user
	}
	c.JSON(http.StatusOK,attentionList{
		Response:model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg: utils.GetStatusMsg(utils.GET_ATTENTION_LIST_SUCCESS),
		},
		UserList: userList,
	})
}

//GetFriendList  获取好友列表
func GetFriendList(c *gin.Context)  {
	userId := c.Query("user_id")
	userIdInt ,_ := strconv.Atoi(userId)
	relation,count := model.GetRelationListWithFollower(int32(userIdInt))
	userList := make([]model.User,count)
	for i:=0;i<int(count);i++{
		user := model.GetUserData(relation[i].FollowingId)
		userList[i]=user
	}
	c.JSON(http.StatusOK,attentionList{
		Response:model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg: utils.GetStatusMsg(utils.GET_ATTENTION_LIST_SUCCESS),
		},
		UserList: userList,
	})
}
