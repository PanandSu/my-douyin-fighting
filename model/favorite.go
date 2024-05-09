package model

import "time"

type Favorite struct {
	FavoriteId uint
	UserId     uint
	VideoId    uint
	CreatedAt  time.Time
	DeletedAt  *time.Time
}
