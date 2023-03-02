package model

import (
	"MiniDouyin/db"
	"gorm.io/gorm"
)

type Relation struct {
	RelationId int32 `json:"relation_id"`
	FollowerId int32 `json:"follower_id"`
	FollowingId int32 `json:"following_id"`
}

//AttentionActionToMysql 关注操作（将关注信息写入数据库）
func AttentionActionToMysql(userId , toUserId int32)  {
	//1.1将关注信息写入关注信息表
	relation := Relation{
		FollowerId: userId,
		FollowingId: toUserId,
	}
	db.Db.Create(&relation)
	//1.2将user_id所对应用户的follow_count + 1
	db.Db.Model(&user{}).Where("id = ?",userId).Update("follow_count",gorm.Expr("follow_count + ?",1))
	//1.3将to_user_id所对应用户的follower_count + 1
	db.Db.Model(&user{}).Where("id = ?",toUserId).Update("follower_count",gorm.Expr("follower_count + ?",1))
}

//DeleteActionToMysql 取消关注（将关注信息写入数据库）
func DeleteActionToMysql(userId,toUserId int32)  {
	db.Db.Where("follower_id = ? && following_id = ?",userId,toUserId).Delete(&Relation{})
}

// GetRelationListWithFollower  根据user_id来获取关注用户id
func GetRelationListWithFollower(userId int32) ([]Relation,int64) {
	var relationList []Relation
	var count int64
	db.Db.Where("follower_id = ?",userId).Find(&relationList).Count(&count)
	return relationList,count
}

// GetRelationListWithFollowing   根据user_id来获取粉丝用户id
func GetRelationListWithFollowing(userId int32) ([]Relation,int64) {
	var relationList []Relation
	var count int64
	db.Db.Where("following_id = ？",userId).Find(&relationList).Count(&count)
	return relationList,count
}