package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	__ "user-api/proto"
	"user-api/request"
)

func UserViewWorks(c *gin.Context) {
	var req request.UserViewWorks
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "未接收参数",
			"data": nil,
		})
	}
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("new grpc client error: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	works, err := c1.UserViewWorks(c, &__.UserViewWorksRequest{UserId: req.UserId})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "作品展示失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "作品展示成功",
		"data": works,
	})
}
