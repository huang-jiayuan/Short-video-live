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
		Source: req.Source,
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
		"code": 10000,
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
			"code": 1000,
			"msg":  "请求失败！无法找到获取的资源",
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
	login, err := c1.UserLogin(c, &__.UserLoginRequest{
		Mobile: req.Mobile,
		Code:   req.Code,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": err.Error(),
		})
		return
	}
	claims := pkg.CustomClaims{ID: uint(login.Greet)}
	token, err := pkg.NewJWT("2211a").CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "服务器内部错误",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "服务器响应正常",
		"data": map[string]interface{}{"token": token},
	})
}

func UserCenter(c *gin.Context) {
	userId := c.GetUint("userId")
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	center, err := c1.UserCenter(c, &__.UserCenterRequest{UserId: int64(userId)})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "服务器内部错误",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "服务器响应正常",
		"data": center,
	})
}

func ImproveUserInfo(c *gin.Context) {
	var req request.ImproveUserInfo
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "请求失败！无法找到获取的资源",
			"data": err.Error(),
		})
		return
	}
	userId := c.GetUint("userId")
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	info, err := c1.ImproveUserInfo(c, &__.ImproveUserInfoRequest{
		Name:          req.Name,
		NickName:      req.NickName,
		UserCode:      req.UserCode,
		SignaTure:     req.SignaTure,
		Sex:           req.Sex,
		Constellation: req.Constellation,
		AvatorFileId:  req.AvatorFileId,
		Mobile:        req.Mobile,
		Age:           req.Age,
		UserId:        int64(userId),
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "服务器响应正常",
		"data": info,
	})
}

func UserLikes(c *gin.Context) {
	var req request.UserLikes
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "请求失败！无法找到获取的资源",
		})
		return
	}
	userId := c.GetUint("userId")
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	likes, err := c1.UserLikes(c, &__.UserLikesRequest{
		UserId: int64(userId),
		WorkId: req.WorkId,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "服务器内部错误",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "服务器响应正常",
		"data": likes,
	})
}
