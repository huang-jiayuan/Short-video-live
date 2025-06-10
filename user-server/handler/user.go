package handler

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"user-server/basic/inits"
	"user-server/models"
	__ "user-server/proto"
)

type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Sendsms(_ context.Context, in *__.SendsmsRequest) (*__.SendsmsResponse, error) {
	code := rand.Intn(9000) + 1000
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(in.Mobile) {
		return nil, errors.New("手机号格式不正确")
	}
	get := inits.RedisGet("SendSmsIncr" + in.Mobile)
	if get == 1 {
		return nil, errors.New("已发送短信，60秒后可重试")
	}

	inits.RedisSet("SendSms"+in.Source+in.Mobile, code)
	incr := inits.RedisIncr("SendSmsIncr" + in.Mobile)
	if incr == 1 {
		inits.RedisExpire("SendSmsIncr" + in.Mobile)
	}
	return &__.SendsmsResponse{Greet: "验证码发送成功"}, nil
}
func (s *Server) UserLogin(_ context.Context, in *__.UserLoginRequest) (*__.UserLoginResponse, error) {
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(in.Mobile) {
		return nil, errors.New("手机号格式不正确")
	}
	get := inits.RedisGet("SendSms" + "login" + in.Mobile)
	if get != int(in.Code) {
		return nil, errors.New("短信验证码错误")
	}
	u := &models.VideoUser{}
	user, err := u.FindUserByMobile(in.Mobile)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	code := rand.Intn(900000) + 100000
	if user.Id == 0 {
		user, err = u.CreateUser(in.Mobile, strconv.Itoa(code))
		if err != nil {
			return nil, errors.New("注册失败")
		}
	}
	fmt.Println(11111)
	users, err := u.FindUserByMobile(in.Mobile)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	fmt.Println(22222)
	return &__.UserLoginResponse{Greet: int64(users.Id)}, nil
}

func (s *Server) UserCenter(_ context.Context, in *__.UserCenterRequest) (*__.UserCenterResponse, error) {
	u := &models.VideoUser{}
	w := &models.VideoWorks{}
	id, err := u.FindUserById(in.UserId)
	if err != nil {
		return nil, errors.New("用户查询失败")
	}
	userId, err := w.FindWorksByUserId(in.UserId)
	if err != nil {
		return nil, errors.New("作品查询失败")
	}
	var Item []*__.UserCenter
	for _, user := range id {
		Item = append(Item, &__.UserCenter{
			Name:          user.Name,
			NickName:      user.NickName,
			UserCode:      user.UserCode,
			Signature:     user.Signature,
			Sex:           user.Sex,
			Constellation: user.Constellation,
			AttendCount:   user.AttendCount,
			FansCount:     user.FansCount,
			ZanCount:      user.ZanCount,
			AvatarFileId:  int64(user.AvatorFileId),
			AuthrityInfo:  user.AuthrityType,
			Mobile:        user.Mobile,
			RealNameAuth:  user.RealNameAuth,
			Age:           int64(user.Age),
			AuthrityType:  user.AuthrityType,
		})
	}
	for _, works := range userId {
		Item = append(Item, &__.UserCenter{
			Title: works.Title,
			Desc:  works.Desc,
		})
	}
	return &__.UserCenterResponse{List: Item}, nil
}

func (s *Server) ImproveUserInfo(_ context.Context, in *__.ImproveUserInfoRequest) (*__.ImproveUserInfoResponse, error) {
	u := &models.VideoUser{}
	err := u.ImproveUserInfo(in)
	if err != nil {
		return nil, errors.New("修改失败")
	}
	return &__.ImproveUserInfoResponse{Greet: "修改成功"}, nil
}

func (s *Server) UserLikes(_ context.Context, in *__.UserLikesRequest) (*__.UserLikesResponse, error) {
	w := &models.VideoWorks{}
	works, err := w.FindVideoWorks(in.WorkId)
	if err != nil {
		return nil, errors.New("作品查询失败")
	}
	if works.Id == 0 {
		return nil, errors.New("作品不存在")
	}
	like := w.LikeCount + 1
	fmt.Println(11111)
	err = w.UpdateLikeCount(in.WorkId, int(like))
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("修改失败")
	}
	fmt.Println(22222)
	return &__.UserLikesResponse{Greet: "点赞成功"}, nil
}
