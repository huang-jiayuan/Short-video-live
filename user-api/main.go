package main

import (
	"github.com/gin-gonic/gin"
	"user-api/pkg"
	"user-api/router"
)

func main() {
	r := gin.Default()
	pkg.InitMinio()
	router.Router(r)
	r.Run(":8080")
}
