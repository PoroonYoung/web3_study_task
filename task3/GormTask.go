package main

import (
	"fmt"
	"gorm.io/gorm"
	db "web3_study_task"
)

type User struct {
	gorm.Model
	Username string
	Posts    []Post
	PostNum  uint
}

func (u *User) TableName() string {
	return "user"
}

type Post struct {
	gorm.Model
	Title        string
	Body         string
	UserID       uint
	User         User
	Comments     []Comment
	CommentState string
}

func (p *Post) TableName() string {
	return "post"
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("post创建前钩子方法进入")
	//1.用户文章数+1
	err = tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_num", gorm.Expr("post_num + 1")).Error
	if err != nil {
		return err
	}
	//2.新文章默认设置评论状态为无评论
	p.CommentState = "无评论"
	fmt.Println("post创建前钩子方法结束")
	return
}

type Comment struct {
	gorm.Model
	Description   string
	PostID        uint
	Post          Post
	PublishUserID uint
	PublishUser   User
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("评论删除后钩子开始", c)
	postID := c.PostID
	//如果删除时的对象中只有id，那这里也只拿得到id，没有别的字段
	if postID == 0 {
		if err := tx.Model(&Comment{}).
			Unscoped(). // 软删后也能查到当前这条
			Select("post_id").
			Where("id = ?", c.ID).
			Scan(&postID).Error; err != nil {
			return err
		}
	}
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", postID).Count(&count)
	if count <= 0 {
		err := tx.Model(&Post{}).Where("id = ?", postID).Update("comment_state", "无评论").Error
		if err != nil {
			return err
		}
	}
	fmt.Println("评论删除后钩子结束")
	return
}

func main() {
	database := db.GetDb()
	database.AutoMigrate(&User{}, &Post{}, &Comment{})
	//已通过AI插入测试数据，详见testInfo.sql

	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息
	//findPostBySomeBodyWithComment(database, "alice")

	//使用Gorm查询评论数量最多的文章信息
	//findMostCommentPost(database)

	//新发布一个文章
	//database.Create(&Post{
	//	Model:        gorm.Model{},
	//	Title:        "新文章标题",
	//	Body:         "文章内容123321312314124",
	//	UserID:       1,
	//	CommentState: "",
	//})

	//删除一个评论
	database.Delete(&Comment{
		Model: gorm.Model{ID: 80},
	})
}

func findMostCommentPost(database *gorm.DB) {
	//select * from post where id = (select post_id from comment group by post_id order by count(*) desc limit 1)
	var post = Post{}
	database.Where("id = (?)", database.Model(&Comment{}).Select("post_id").Group("post_id").Order("count(*) desc").Limit(1)).First(&post)
	fmt.Println(post)
}

func findPostBySomeBodyWithComment(database *gorm.DB, name string) {
	var User User
	database.Preload("Posts").Preload("Posts.Comments").Preload("Posts.Comments.PublishUser").First(&User, "username = ?", name)
	fmt.Println(User.Username)
	for _, post := range User.Posts {
		fmt.Println("这是文章标题", post.Title, "文章内容", post.Body)
		for _, comment := range post.Comments {
			fmt.Println("这是评论", comment.Description, "发布者", comment.PublishUser.Username)
		}
		fmt.Println("一篇文章结束")
	}
}
