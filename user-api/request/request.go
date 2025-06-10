package request

// TODO 短信验证码请求字段
type Sendsms struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Score  string `form:"score" json:"score" binding:"required"`
}

// TODO 用户登录请求字段
type UserLogin struct {
	Mobile  string `form:"mobile" json:"mobile" binding:"required"`
	Sendsms string `form:"sendsms" json:"sendsms" binding:"required"`
}

// TODO 用户个人首页

type UserHomepage struct {
	Id int64 `form:"id" json:"id" binding:"required"`
}

// TODO  查看其他用户的作品
type UserViewWorks struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required"`
}
