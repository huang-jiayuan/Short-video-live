package router

import (
	"github.com/gin-gonic/gin"
	"user-api/handler"
	"user-api/pkg"
)

func Router(r *gin.Engine) {
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/sendsms", handler.Sendsms)
			user.POST("/login", handler.Login)
			user.Use(pkg.JWTAuth("2211a"))
			user.POST("/center", handler.UserCenter)
			user.POST("/improve/info", handler.ImproveUserInfo)
			user.POST("/like", handler.UserLikes)
		}
		work := api.Group("/work")
		{
			work.POST("/video/list", handler.VideoWorksList)
			work.Use(pkg.JWTAuth("2211a"))
			work.POST("/add", handler.AddVideoWorks)
			work.POST("/add/comment", handler.CreateVideoWorksComment)
		}
	}
}
