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
			user.GET("/homepage", handler.UserHomepage)
		}
		video := api.Group("/video")
		{
			video.GET("/viewworks", handler.UserViewWorks)
			video.Use(pkg.JWTAuth("2211a"))
		}
	}
}
