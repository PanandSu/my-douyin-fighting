package service

import (
	"my-douyin-fighting/model"
)

func GetLikeStatus(uid, vid uint) (bool, error) {

}
func Like(uid, vid uint) error {

}
func Unlike(uid, vid uint) error {

}
func GetLikedVids(uid uint) ([]uint, error) {

}
func GetLikedVideos(uid uint) ([]model.Video, error) {

}
func GetLikeStatuses(uid uint, vids []uint) ([]bool, error) {
}

func GetLikeCount(vids []uint) ([]uint, error) {

}
