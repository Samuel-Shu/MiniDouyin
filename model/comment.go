package model

import (
	"MiniDouyin/db"
	"gorm.io/gorm"
	"time"
)

//PublishCommentToMysql 将评论内容插入数据库
func PublishCommentToMysql(userId , videoId int32, content string) Comment {
	comment := Comment{
		Id:userId,
		VideoId: videoId,
		Content: content,
		CreateTime: time.Now().Format("01-02"),
	}
	commentRes := Comment{}
	db.Db.Create(&comment).Select("comment_id","id","content","create_time").Scan(&commentRes)
	db.Db.Model(&video{}).Where("video_id = ?",videoId).Update("comment_count",gorm.Expr("comment_count + ?",1))
	return commentRes
}

//DeleteCommentWithCommentId 依照commentId来删除评论
func DeleteCommentWithCommentId(commentId ,videoId int32)  {
	db.Db.Where("comment_id = ?",commentId).Delete(&Comment{})
	db.Db.Model(&video{}).Where("video_id = ?",videoId).Update("comment_count",gorm.Expr("comment_count - ?",1))
}


//GetCommentList 根据video_id来获取评论列表
func GetCommentList(videoId int32) ([]Comment,int64) {
	var commentList []Comment
	var count int64
	db.Db.Where("video_id = ?",videoId).Find(&commentList).Count(&count)
	return commentList,count
}