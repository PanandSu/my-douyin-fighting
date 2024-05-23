package model

import "time"

type User struct {
	Id            uint
	Name          string
	Password      string
	FollowCount   uint
	FollowerCount uint
	TotalLike     uint
	LikeCount     uint
	CreatedAt     time.Time
	ExtInfo       *string
}
