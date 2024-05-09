package model

import "time"

type Follow struct {
	FollowId    uint
	CelebrityId uint
	FollowerId  uint
	IsFollow    bool
	CreatedAt   time.Time
	DeletedAt   *time.Time
}
