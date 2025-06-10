package handler

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"time"
	"user-server/basic/inits"
	"user-server/models"
	__ "user-server/proto"
)

type Server struct {
	__.UnimplementedUserServer
}

// 获取验证码
func (s *Server) Sendsms(_ context.Context, in *__.SendsmsRequest) (*__.SendsmsResponse, error) {
	one := inits.RedisClient.Get(context.Background(), "Send"+in.Score+in.Mobile).Val()
	if one >= "1" {
		return nil, errors.New("每分钟只能发送一条消息")
	}
	code := rand.Intn(9000) + 1000
	inits.RedisClient.Set(context.Background(), "SendOne"+in.Score+in.Mobile, "1", time.Minute*1)
	inits.RedisClient.Set(context.Background(), "Send"+in.Score+in.Mobile, code, time.Minute*5)
	return &__.SendsmsResponse{
		Greet: "验证码发送成功",
		Code:  0,
	}, nil
}

// 用户登录注册一体化
func (s *Server) Login(_ context.Context, in *__.LoginRequest) (*__.LoginResponse, error) {
	user := models.VideoUser{}
	number := user.ValidatePhoneNumber(in.Mobile)
	if !number {
		fmt.Println("手机号格式不正确")
		return nil, errors.New("手机号格式不正确")
	}
	if in.Mobile == "" {
		fmt.Println("手机号不能为空")
		return nil, errors.New("手机号不能为空")
	}
	get := inits.RedisClient.Get(context.Background(), "SendLogin"+in.Mobile).Val()
	if get != in.Sendsms {
		fmt.Println("短信验证码验证错误")
		return nil, errors.New("短信验证码验证错误")
	}
	if get == "" {
		fmt.Println("验证码已经过期")
		return nil, errors.New("验证码已经过期")
	}
	err := user.Find(in.Mobile)
	if err != nil {
		fmt.Println("未查询到此用户的存在")
	}
	if user.Id == 0 {
		fmt.Println("当前用户未注册立即注册")
		err = user.Create(in)
		if err != nil {
			fmt.Println("用户注册失败")
			return nil, err
		}
	}
	inits.RedisClient.Del(context.Background(), "SendLogin"+in.Mobile)
	return &__.LoginResponse{
		Id:    int64(user.Id),
		Greet: "登录成功",
	}, nil

}

// 用户首页
func (s *Server) UserHomepage(_ context.Context, in *__.UserHomepageRequest) (*__.UserHomepageResponse, error) {
	user := models.VideoUser{}
	err := user.GetUserById(in.Id)
	if err != nil {
		fmt.Println("没有当前的用户")
		return nil, err
	}
	li, err := user.ListUser(in.Id)
	if err != nil {
		return nil, err
	}
	var list []*__.UserItem
	for _, video := range li {
		lists := &__.UserItem{
			Name:          video.Name,
			NickName:      video.NickName,
			UserCode:      video.UserCode,
			Signature:     video.Signature,
			Sex:           video.Sex,
			IpAddress:     video.IpAddress,
			Constellation: video.Constellation,
			AttendCount:   video.AttendCount,
			FansCount:     video.FansCount,
			ZanCount:      video.ZanCount,
			AvatorFileId:  int64(video.AvatorFileId),
			AuthriryInfo:  video.AuthriryInfo,
			RealNameAuth:  video.RealNameAuth,
		}
		list = append(list, lists)
	}
	return &__.UserHomepageResponse{List: list}, nil
}

// 查看其他用户的作品
func (s *Server) UserViewWorks(_ context.Context, in *__.UserViewWorksRequest) (*__.UserViewWorksResponse, error) {
	works := models.VideoWorks{}
	err := works.GetVideoWorksById(in.UserId)
	if err != nil {
		fmt.Println("没有查到此用户")
		return nil, err
	}
	li, err := works.GetListVideoWorksById(in.UserId)
	if err != nil {
		fmt.Println("作品展示失败")
		return nil, err
	}
	var list []*__.VideoWorksItem
	for _, video := range li {
		lists := &__.VideoWorksItem{
			Title:          video.Title,
			Desc:           video.Desc,
			MusicId:        int64(video.MusicId),
			WorkType:       video.WorkType,
			CheckStatus:    video.CheckStatus,
			CheckUser:      int64(video.CheckUser),
			IpAddress:      video.IpAddress,
			WorkPermission: video.WorkPermission,
			LikeCount:      int64(video.LikeCount),
			CommentCount:   int64(video.CommentCount),
			ShareCount:     int64(video.ShareCount),
			CollectCount:   int64(video.CollectCount),
			BrowseCount:    int64(video.BrowseCount),
			TopicId:        int64(video.TopicId),
		}
		list = append(list, lists)
	}
	return &__.UserViewWorksResponse{List: list}, nil
}
