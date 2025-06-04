package router

import (
	"github.com/gin-gonic/gin"
	"user-api/handler"
)

func Router(r *gin.Engine) {
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/sendsms", handler.Sendsms)
		}
	}
}
