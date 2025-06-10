package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"user-api/pkg"
	__ "user-api/proto"
	"user-api/request"
)

func Sendsms(c *gin.Context) {
	var req request.Sendsms
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	sendsms, err := c1.Sendsms(c, &__.SendsmsRequest{
		Mobile: req.Mobile,
		Score:  req.Score,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "短信发送失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "短信发送成功",
		"data": sendsms,
	})
	return
}

func Login(c *gin.Context) {
	var req request.UserLogin
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "未接收任何参数",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	login, err := c1.Login(c, &__.LoginRequest{
		Mobile:  req.Mobile,
		Sendsms: req.Sendsms,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "登录失败",
			"data": err.Error(),
		})
		return
	}
	token, err := pkg.NewJWT("2211a").CreateToken(pkg.CustomClaims{
		ID: uint(login.Id)})
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": token,
	})
}

func UserHomepage(c *gin.Context) {
	var req request.UserHomepage
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "未接收到参数",
			"data": nil,
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	homepage, err := c1.UserHomepage(c, &__.UserHomepageRequest{Id: req.Id})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "用户个人展示失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户个人展示成功",
		"data": homepage,
	})
}
