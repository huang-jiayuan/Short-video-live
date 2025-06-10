package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"path/filepath"
	"user-api/pkg"
	__ "user-api/proto"
	"user-api/request"
)

func VideoWorksList(c *gin.Context) {
	var req request.VideoWorksList
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
	list, err := c1.VideoWorksList(c, &__.VideoWorksListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
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
		"data": list,
	})
}
func AddVideoWorks(c *gin.Context) {
	var req request.AddVideoWorks
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
	userId := c.GetUint("userId")
	ip := c.ClientIP()
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "请求失败！无法找到获取的资源",
			"data": err.Error(),
		})
		return
	}
	fileHeader, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer fileHeader.Close()
	ext := filepath.Ext(file.Filename)
	if ext != ".mp4" && ext != ".jpg" && ext != ".png" && ext != ".webp" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": "文件格式错误",
		})
		return
	}

	if file.Size >= 500*1024*1024 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": "文件过大",
		})
		return
	}

	err = pkg.Upload(file.Filename, fileHeader, file.Size)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": "文件上传失败",
		})
		return
	}

	works, err := c1.AddVideoWorks(c, &__.AddVideoWorksRequest{
		Title:          req.Title,
		Desc:           req.Desc,
		MusicId:        req.MusicId,
		WorkType:       req.WorkType,
		WorkPermission: req.WorkPermission,
		UserId:         int64(userId),
		IpAddress:      ip,
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
		"data": works,
	})

}

func CreateVideoWorksComment(c *gin.Context) {
	var req request.CreateVideoWorksComment
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
	userId := c.GetUint("userId")
	comment, err := c1.CreateVideoWorksComment(c, &__.CreateVideoWorksCommentRequest{
		WorkId:  req.WorkId,
		UserId:  int64(userId),
		Content: req.Content,
		Tag:     req.Tag,
		Pid:     req.Pid,
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
		"data": comment,
	})
}
