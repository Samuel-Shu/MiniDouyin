package model

import (
	"miniVersionDouyin/db"
)

type UserRegister struct {
	Status Response
	UserId int32  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type user struct {
	Id       int32
	Username string
	Password string
}

// Register 用户注册
func Register(username, password string) int32 {
	if FindUser(username) == 0 {
		db.Db.Model(&user{}).Create(map[string]interface{}{"username": username, "password": password})
		return 0
	}
	return 1
}

// FindUser 查找用户并返回其id
func FindUser(username string) int32 {
	user := user{}
	db.Db.Where("username=?", username).First(&user)
	return user.Id
}

func GetUserData(id int32) User {
	user := User{}
	db.Db.Where("id=?", id).First(&user)
	return user
}

func Login(username, password string) bool {
	userData := user{}
	var res bool
	if FindUser(username) == 0 {
		res = false
	} else {
		db.Db.Where("username=?", username).First(&userData)
		if userData.Password == password {
			res = true
		} else {
			res = false
		}
	}
	return res
}
