package handler

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
	r.GET("ping", Ping)

	// user
	r.POST("/douyin/user/register/", RegisterHandler)
	r.POST("/douyin/user/login/", LoginHandler)
	r.GET("/douyin/user/", GetUserInfoHandler)

	// video
	r.POST("/douyin/publish/action/", VideoPublishHandler)
	r.GET("/douyin/publish/list/", VideoPublishedListHandler)
	r.GET("/douyin/feed/", VideoFeedHandler)

	// comment
	r.POST("/douyin/comment/action/", CommentHandler)
	r.GET("/douyin/comment/list/", CommentListHandler)
}
