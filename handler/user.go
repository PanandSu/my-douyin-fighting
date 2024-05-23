package handler

import (
	"github.com/gin-gonic/gin"
	gb "my-douyin-fighting/glob"
	"my-douyin-fighting/service"
	"my-douyin-fighting/util"
	"regexp"
	"strconv"
	"unicode/utf8"
)

type UserResponse struct {
	Response
	User User
}

type UserLoginResponse struct {
	Response
	Uid   uint
	Token string
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	length := utf8.RuneCountInString(username)
	if length <= 0 || length > gb.MaxUsernameLength {
		c.JSON(200, Response{
			Code: 0,
			Msg:  "",
		})
		return
	}
	if match, _ := regexp.MatchString(gb.PasswordPattern, password); !match {
		c.JSON(200, Response{
			Code: 0,
			Msg:  "",
		})
		return
	}
	user, err := service.Register(username, password)
	if err != nil {
		c.JSON(200, Response{
			Code: 0,
			Msg:  err.Error(),
		})
		return
	}
	token, err := util.GenerateToken(user)
	if err != nil {
		c.JSON(200, Response{
			Code: 500,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, UserLoginResponse{
		Response: Response{
			Code: 0,
			Msg:  "",
		},
		Uid:   user.Id,
		Token: token,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := service.Login(username, password)
	if err != nil {
		c.JSON(200, Response{
			Code: 0,
			Msg:  err.Error(),
		})
		return
	}
	token, err := util.GenerateToken(user)
	if err != nil {
		c.JSON(500, Response{
			Code: 1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, UserLoginResponse{
		Response: Response{
			Code: 0,
			Msg:  "",
		},
		Uid:   user.Id,
		Token: token,
	})
}

func UserInfo(c *gin.Context) {
	userId := c.Query("user_id")
	uid, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		c.JSON(200, Response{
			Code: 0,
			Msg:  err.Error(),
		})
		return
	}
	user, err := service.GetUserInfo(uint(uid))
	if err != nil {
		c.JSON(500, Response{
			Code: 0,
			Msg:  err.Error(),
		})
	}
	viewerId := c.GetUint64("UserId")
	followed, err := service.GetFollowStatus(uint(uid), uint(viewerId))
	if err != nil {
		c.JSON(500, Response{
			Code: 0,
			Msg:  err.Error(),
		})
	}
	c.JSON(200, UserResponse{
		Response: Response{
			Code: 0,
			Msg:  "",
		},
		User: User{
			ID:            uint(uid),
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			TotalLike:     user.TotalLike,
			LikeCount:     user.LikeCount,
			IsFollow:      followed,
		},
	})
}
