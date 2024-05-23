package handler

type User struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint   `json:"follow_count,omitempty"`
	FollowerCount uint   `json:"follower_count,omitempty"`
	TotalLike     uint   `json:"total_favorited,omitempty"`
	LikeCount     uint   `json:"favorite_count,omitempty"`
	IsFollow      bool   `json:"is_follow"`
}

type Video struct {
}
type Comment struct {
}

type Response struct {
	Code int
	Msg  string
}
