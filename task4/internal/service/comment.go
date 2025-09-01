package service

import (
	"web3_study_task/task4/internal/repository"
	"web3_study_task/task4/pkg/dto"
	"web3_study_task/task4/pkg/entity"
)

type CommentService struct {
	commentRepo *repository.CommentRepository
}

func NewCommentService(repo *repository.CommentRepository) *CommentService {
	return &CommentService{commentRepo: repo}
}

func (s CommentService) GetAllByPostId(postId uint) (*[]entity.Comment, error) {
	return s.commentRepo.GetAllByPostId(postId)
}

func (s CommentService) PublishComment(req *dto.CommentPublishRequest) error {
	comment := &entity.Comment{
		Content: req.Content,
		PostID:  req.PostID,
		UserID:  req.UserID,
	}
	return s.commentRepo.PublishComment(comment)
}
