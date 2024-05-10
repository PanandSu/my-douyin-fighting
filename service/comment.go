package service

import (
	"fmt"
	"gorm.io/gorm"
	gb "my-douyin-fighting/glob"
	"my-douyin-fighting/model"
)

func AddComment(comment *model.Comment) error {
	gb.DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(comment)
		AddCommentRedis(comment)
	})
}

func DelComment(uid, vid, cid uint) error {
	var comment *model.Comment
	comment.Id = cid
	gb.DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("user_id=? and video_id=?", uid, vid).Delete(comment)
		DelCommentRedis(vid, cid)
	})
}

func GetUserList(vid uint) (*[]model.User, error) {
	key := fmt.Sprintf("CommentsOfVideo:%d", vid)
	var users *[]model.User
}

func GetCommentList(vid uint) (*[]model.Comment, error) {
	var comments *[]model.Comment
}

func GetCommentCountList(vids []uint) (*[]uint, error) {

}

func GetCommentCountList2(vids []uint) (*[]uint, error) {

}
