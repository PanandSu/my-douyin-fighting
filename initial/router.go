package initial

import (
	"github.com/gin-gonic/gin"
	"my-douyin-fighting/glob"
	"my-douyin-fighting/handler"
	"strconv"
)

func Route() {
	r := gin.Default()
	r.Static("/static", "./public")
	api := r.Group("/douyin")
	{
		api.GET("/feed", handler.Feed)
		api.POST("/user/register", handler.Register)
		api.POST("/user/login", handler.Login)
		api.GET("/publish/list", handler.PublishList)

		api.GET("/favorite/list", handler.FavoriteList)
		api.GET("/comment/list", handler.CommentList)

		api.GET("/relation/follow/list", handler.FollowList)
		api.GET("/relation/follower/list", handler.FollowerList)
	}
	addr := glob.Config.GinConfig.Host + strconv.Itoa(glob.Config.GinConfig.Port)
	if err := r.Run(addr); err != nil {
		panic(err)
	}
}
