package model

import "time"

type User struct {
	Id            uint
	Name          string
	Password      string
	FollowCount   uint
	FollowerCount uint
	CreatedAt     time.Time
	ExtInfo       *string
	TotalLike     int
	LikeCount     int
}
