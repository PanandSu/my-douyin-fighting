package handler

import (
	"github.com/gin-gonic/gin"
	"my-douyin-fighting/service"
	"strconv"
)

type UserListResponse struct {
	Response
	Users []User `json:"users"`
}

func RelationAction(c *gin.Context) {
	userId := c.Param("userid")
	action := c.Param("action")
	act, err := strconv.Atoi(action)
	if err != nil {
		c.JSON(400, Response{
			Code: 1,
			Msg:  err.Error(),
		})
		return
	}
	if act != 1 && act != 2 {
		c.JSON(400, Response{
			Code: 1,
			Msg:  err.Error(),
		})
		return
	}
	viewerId := c.GetUint("UserId")
	if act == 1 {
		err := service.Follow(viewerId, userId)
		if err != nil {
			c.JSON(200, Response{
				Code: 0,
				Msg:  "success",
			})
			return
		}
	} else {
		err := service.Unfollow(viewerId, userId)
		if err != nil {
			c.JSON(200, Response{
				Code: 0,
				Msg:  "success",
			})
			return
		}
	}
	c.JSON(200, Response{
		Code: 0,
		Msg:  "success",
	})
}
func Follow(c *gin.Context)   {}
func Unfollow(c *gin.Context) {}

func FollowList(c *gin.Context)   {}
func FollowerList(c *gin.Context) {}
