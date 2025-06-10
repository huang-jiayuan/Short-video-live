package request

type Sendsms struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Source string `form:"source" json:"source" binding:"required"`
}
type UserLogin struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
	Code   int64  `json:"code" form:"code" binding:"required"`
}

type ImproveUserInfo struct {
	Name          string `json:"name" form:"name"`
	NickName      string `json:"nick_name" form:"nick_name"`
	UserCode      string `json:"user_code" form:"user_code"`
	SignaTure     string `json:"signa_ture" form:"signa_ture"`
	Sex           string `json:"sex" form:"sex"`
	Constellation string `json:"constellation" form:"constellation"`
	AvatorFileId  int64  `json:"avator_file_id" form:"avator_file_id"`
	Mobile        string `json:"mobile" form:"mobile"`
	Age           int64  `json:"age" form:"age"`
}

type VideoWorksList struct {
	Page     int64 `json:"page" form:"page"`
	PageSize int64 `json:"pageSize" json:"pageSize"`
}

type AddVideoWorks struct {
	Title          string `form:"title" json:"title"`                     //标题
	Desc           string `form:"desc" json:"desc"`                       //描述
	MusicId        int64  `form:"music_id" json:"music_id"`               //选择音乐
	WorkType       string `form:"work_type" json:"work_type"`             //作品类型
	WorkPermission string `form:"work_permission" json:"work_permission"` //作品权限
}
type CreateVideoWorksComment struct {
	WorkId  int64  `form:"work_id" json:"work_id"` //作品id
	Content string `form:"content" json:"content"` //评论内容
	Tag     int64  `form:"tag" json:"tag"`         //评论标签表
	Pid     int64  `form:"pid" json:"pid"`         //父级id
}
type UserLikes struct {
	WorkId int64 `form:"work_id" json:"work_id"`
}
