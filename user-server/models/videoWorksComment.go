package models

type VideoWorkComment struct {
	Id      int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	WorkId  int32  `gorm:"column:work_id;type:int;comment:作品ID;default:NULL;" json:"work_id"`            // 作品ID
	UserId  int32  `gorm:"column:user_id;type:int;comment:用户ID;default:NULL;" json:"user_id"`            // 用户ID
	Content string `gorm:"column:content;type:varchar(100);comment:评论内容;default:NULL;" json:"content"` // 评论内容
	Tag     int32  `gorm:"column:tag;type:int;comment:评论标签表;default:NULL;" json:"tag"`                // 评论标签表
	Pid     int32  `gorm:"column:pid;type:int;comment:父级ID;default:0;" json:"pid"`                       // 父级ID
}
