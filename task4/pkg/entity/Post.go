package entity

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint      `gorm:"not null"`                        // 外键字段
	User     User      `gorm:"foreignKey:UserID;references:ID"` // 多对一关联：帖子的作者
	Comments []Comment `gorm:"foreignKey:PostID;references:ID"` // 一对多关联：帖子的评论
}

func (p *Post) TableName() string {
	return "post"
}
