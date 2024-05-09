package model

import "time"

type Video struct {
	VideoId       uint
	Title         string
	PlayName      string
	CoverName     string
	CommentCount  uint
	FavoriteCount uint
	AuthorId      uint
	CreatedAt     time.Time
	ExtInfo       *string
}

type VideoCount struct {
	VideoId      uint
	CommentCount uint
}
