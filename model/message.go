package model

import (
	"MiniDouyin/db"
	"time"
)

func SendMessage(userId , toUserId int32,content string)  {
	message := Message{
		ToUserId:toUserId,
		FromUserId: userId,
		Content: content,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	db.Db.Create(&message)
}

func GetMessage(userId , toUserId int32) []Message {
	var messageList []Message
	db.Db.Where("from_user_id = ? && to_user_id = ?",userId,toUserId).Find(&messageList)
	return messageList
}