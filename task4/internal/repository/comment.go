package repository

import (
	"gorm.io/gorm"
	db "web3_study_task"
	"web3_study_task/task4/pkg/entity"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r CommentRepository) GetAllByPostId(postId uint) (*[]entity.Comment, error) {
	commentList := &[]entity.Comment{}
	err := db.DB.Where("post_id = ?", postId).Find(commentList).Error
	return commentList, err
}

func (r CommentRepository) PublishComment(comment *entity.Comment) error {
	return db.DB.Create(comment).Error
}
