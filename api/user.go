package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniVersionDouyin/model"
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
	if code == 0 {
		status.Status.StatusCode = code
		status.Status.StatusMsg = "Register Successfully"
		status.UserId = model.FindUser(username)
		status.Token = "qwertyui"
		fmt.Println("注册成功")
		c.JSON(http.StatusOK, status)
	} else {
		status.Status.StatusCode = 1
		status.Status.StatusMsg = "Register False"
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
	userMsg.Status.StatusCode = 0
	userMsg.Status.StatusMsg = "success"
	c.JSON(http.StatusOK, userMsg)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := model.UserRegister{}
	res := model.Login(username, password)
	if res {
		user.Status.StatusCode = 0
		user.Status.StatusMsg = "success"
		user.UserId = model.FindUser(username)
		user.Token = "fsfsksdfjk"
	} else {
		user.Status.StatusCode = 1
		user.Status.StatusMsg = "false"
	}
	c.JSON(http.StatusOK, user)
}
