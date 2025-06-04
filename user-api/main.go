package main

import (
	"github.com/gin-gonic/gin"
	"user-api/router"
)

func main() {
	r := gin.Default()
	router.Router(r)
	r.Run(":8080")
}
