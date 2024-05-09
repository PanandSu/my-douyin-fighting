package model

import "time"

func init() {

}

type Comment struct {
	CommentId uint   `gorm:"primary_key"`
	UserId    uint   `gorm:"index"`
	VideoId   uint   `gorm:"column:video_id"`
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	DeletedAt *time.Time
}
