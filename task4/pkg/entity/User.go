package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"unique"`
	Password string    `gorm:"size:255" json:"-"`
	Email    string    `gorm:"size:255"`
	Posts    []Post    `gorm:"foreignKey:UserID;references:ID"` // 一对多关联：用户的帖子
	Comments []Comment `gorm:"foreignKey:UserID;references:ID"` // 一对多关联：用户的评论
}

func (u *User) TableName() string {
	return "user"
}
