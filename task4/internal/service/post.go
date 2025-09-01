package service

import (
	"fmt"
	"web3_study_task/task4/internal/repository"
	"web3_study_task/task4/pkg/dto"
	"web3_study_task/task4/pkg/entity"
)

type PostService struct {
	postRepo *repository.PostRepository
	userRepo *repository.UserRepository
}

func NewPostService(postRepo *repository.PostRepository, userRepo *repository.UserRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

// CreatePost 创建文章
func (s *PostService) CreatePost(userID uint, req *dto.CreatePostRequest) (*dto.PostResponse, error) {
	// 验证用户是否存在
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}

	// 创建文章实体
	post := &entity.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	// 保存到数据库
	if err := s.postRepo.Create(post); err != nil {
		return nil, fmt.Errorf("创建文章失败: %v", err)
	}

	// 返回响应
	return &dto.PostResponse{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		UserID:   post.UserID,
		Username: user.Username,
		CreateAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetPost 获取文章详情
func (s *PostService) GetPost(id uint) (*dto.PostResponse, error) {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("文章不存在")
	}

	return &dto.PostResponse{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		UserID:   post.UserID,
		Username: post.User.Username,
		CreateAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetUserPosts 获取用户的文章列表
func (s *PostService) GetUserPosts(userID uint, page, pageSize int) (*dto.PostListResponse, error) {
	offset := (page - 1) * pageSize
	posts, total, err := s.postRepo.GetByUserID(userID, offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("获取文章列表失败: %v", err)
	}

	var postResponses []dto.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, dto.PostResponse{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			UserID:   post.UserID,
			Username: post.User.Username,
			CreateAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &dto.PostListResponse{
		Posts: postResponses,
		Total: total,
	}, nil
}

// GetAllPosts 获取所有文章列表
func (s *PostService) GetAllPosts(page, pageSize int) (*dto.PostListResponse, error) {
	offset := (page - 1) * pageSize
	posts, total, err := s.postRepo.GetAll(offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("获取文章列表失败: %v", err)
	}

	var postResponses []dto.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, dto.PostResponse{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			UserID:   post.UserID,
			Username: post.User.Username,
			CreateAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &dto.PostListResponse{
		Posts: postResponses,
		Total: total,
	}, nil
}

// UpdatePost 根据id+UserId修改文章
func (s *PostService) UpdatePost(req dto.UpdatePostRequest) (int64, error) {
	// 调用repository更新文章
	rowsAffected, err := s.postRepo.UpdateByIdAndUserId(req.ID, req.UserID, req.Title, req.Content)
	if err != nil {
		return 0, fmt.Errorf("更新文章失败: %v", err)
	}

	// 如果没有更新任何行，说明文章不存在或不属于该用户
	if rowsAffected == 0 {
		return 0, fmt.Errorf("文章不存在或无权限修改")
	}

	return rowsAffected, nil
}

func (s *PostService) DeletePost(id uint, userId uint) error {
	//调用repository删除文章
	rowsAffected, err := s.postRepo.Delete(id, userId)
	if err != nil {
		return fmt.Errorf("删除文章失败: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("文章不存在或无权限删除")
	}
	return nil
}
