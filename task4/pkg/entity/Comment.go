package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID  uint `gorm:"not null"`                        // 外键字段
	Post    Post `gorm:"foreignKey:PostID;references:ID"` // 多对一关联：评论所属的帖子
	UserID  uint `gorm:"not null"`                        // 外键字段
	User    User `gorm:"foreignKey:UserID;references:ID"` // 多对一关联：评论的作者
	Content string
}

func (c *Comment) TableName() string {
	return "comment"
}
