package service

import "my-douyin-fighting/model"

type likeREdis struct{}

func (r *likeREdis) GetLikeStatus(uid, vid uint) (bool, error) {

}
func (r *likeREdis) AddLikeVids(uid uint, likes []model.Like) {

}
func (r *likeREdis) Like() {

}
func (r *likeREdis) Unlike() {

}
func (r *likeREdis) GetLikeVids() {

}
func (r *likeREdis) GetLikeCount() {

}
func (r *likeREdis) AddLikeCount() {

}
func (r *likeREdis) GetLikeCountList() {

}
func (r *likeREdis) AddLikeCountList() {

}
