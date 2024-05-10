package model

import "time"

type Like struct {
	Id        uint
	UserId    uint
	VideoId   uint
	IsLike    bool
	CreatedAt time.Time
	DeletedAt *time.Time
}
