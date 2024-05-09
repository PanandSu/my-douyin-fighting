package model

import "time"

type User struct {
	UserId        uint
	Name          string
	Password      string
	FollowCount   uint
	FollowerCount uint
	CreatedAt     time.Time
	ExtInfo       *string
}
