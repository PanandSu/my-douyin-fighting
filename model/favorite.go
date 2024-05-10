package model

import "time"

type Like struct {
	Id        uint
	UserId    uint
	VideoId   uint
	CreatedAt time.Time
	DeletedAt *time.Time
}
