package request

type Sendsms struct {
	Mobile string `gorm:"mobile" json:"mobile" binding:"required"`
	Score  string `gorm:"score" json:"score" binding:"required"`
}
