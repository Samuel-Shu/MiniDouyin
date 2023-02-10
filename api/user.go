package api

import (
	"MiniDouyin/model"
	"MiniDouyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserMsg struct {
	Status model.Response
	user   model.User
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	status := model.UserRegister{}
	code := model.Register(username, password)
	if code == utils.SUCCESS {
		status.Status.StatusCode = code
		status.Status.StatusMsg = utils.GetStatusMsg(utils.USER_SUCCESS_REGISTER)
		status.UserId = model.FindUser(username)
		status.Token = "qwertyui"
		fmt.Println("注册成功")
		c.JSON(http.StatusOK, status)
	} else {
		status.Status.StatusCode = utils.FAIL
		status.Status.StatusMsg = utils.GetStatusMsg(utils.USER_FAIL_REGISTER)
		fmt.Println("注册失败")
		fmt.Println(status)
		c.JSON(http.StatusOK, status)
	}
}

func GetUserData(c *gin.Context) {
	userMsg := UserMsg{}
	userId := c.Query("user_id")
	_ = c.Query("token")
	userid, _ := strconv.Atoi(userId)
	user := model.GetUserData(int32(userid))
	userMsg.user = user
	userMsg.Status.StatusCode = utils.SUCCESS
	userMsg.Status.StatusMsg = utils.GetStatusMsg(utils.SUCCESS)
	c.JSON(http.StatusOK, userMsg)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := model.UserRegister{}
	res := model.Login(username, password)
	if res {
		user.Status.StatusCode = utils.SUCCESS
		user.Status.StatusMsg = utils.GetStatusMsg(utils.USER_SUCCESS_LOGIN)
		user.UserId = model.FindUser(username)
		user.Token = "fsfsksdfjk"
	} else {
		user.Status.StatusCode = utils.FAIL
		user.Status.StatusMsg = utils.GetStatusMsg(utils.USER_FAIL_LOGIN)
	}
	c.JSON(http.StatusOK, user)
}
