package dao

type PullEvent struct {
	Id      int64  `gorm:"primaryKey,autoIncrement"`
	Uid     int64  `gorm:"column:uid;index"`
	Type    string `gorm:"column:type"`
	Content string `gorm:"column:content"`
	Ctime   int64  `gorm:"index"`
}
