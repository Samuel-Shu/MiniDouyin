package api

import (
	"MiniDouyin/middleware"
	"MiniDouyin/model"
	"MiniDouyin/rpc/rpcClient"
	pb "MiniDouyin/rpc/rpcpb"
	"MiniDouyin/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserMsg struct {
	model.Response
	User model.User `json:"user"`
}

// Register 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	md5Req := &pb.Md5Req{Password: password}
	md5Password, err := rpcClient.Client.Md5(context.Background(), md5Req)
	if err != nil {
		log.Fatal(err)
	}
	code := model.Register(username, md5Password.GetMd5Password())
	if code == utils.SUCCESS {
		status := model.UserRegister{
			Response: model.Response{
				StatusCode: code,
				StatusMsg:  utils.GetStatusMsg(utils.USER_SUCCESS_REGISTER),
			},
			UserId: model.FindUser(username),
			Token:  middleware.GenerateToken(username, model.FindUser(username)),
		}
		fmt.Println("注册成功")
		c.JSON(http.StatusOK, status)
	} else {
		status := model.UserRegister{
			Response: model.Response{
				StatusCode: utils.FAIL,
				StatusMsg:  utils.GetStatusMsg(utils.USER_FAIL_REGISTER),
			},
		}
		fmt.Println("注册失败")
		fmt.Println(status)
		c.JSON(http.StatusOK, status)
	}
}

// GetUserData 获取用户信息
func GetUserData(c *gin.Context) {
	userId := c.Query("user_id")
	userid, _ := strconv.Atoi(userId)
	userMsg := UserMsg{
		Response: model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg:  utils.GetStatusMsg(utils.SUCCESS),
		},
		User: model.GetUserData(int32(userid)),
	}
	c.JSON(http.StatusOK, userMsg)
}

//Login 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	md5Req := &pb.Md5Req{Password: password}
	md5Password, err := rpcClient.Client.Md5(context.Background(), md5Req)
	if err != nil {
		log.Fatal(err)
	}
	res := model.Login(username, md5Password.GetMd5Password())
	if res {
		user := model.UserRegister{
			Response: model.Response{
				StatusCode: utils.SUCCESS,
				StatusMsg:  utils.GetStatusMsg(utils.USER_SUCCESS_LOGIN),
			},
			UserId: model.FindUser(username),
			Token:  middleware.GenerateToken(username, model.FindUser(username)),
		}
		c.JSON(http.StatusOK, user)
	} else {
		user := model.UserRegister{
			Response: model.Response{
				StatusCode: utils.FAIL,
				StatusMsg:  utils.GetStatusMsg(utils.USER_PASSWORD_IS_NOT_CORRECT),
			},
			UserId: 0,
		}
		c.JSON(http.StatusOK, user)
	}
}
