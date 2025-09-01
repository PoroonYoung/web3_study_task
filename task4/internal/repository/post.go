package repository

import (
	"gorm.io/gorm"
	"web3_study_task/task4/pkg/entity"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

// Create 创建文章
func (r *PostRepository) Create(post *entity.Post) error {
	return r.db.Create(post).Error
}

// GetByID 根据ID获取文章
func (r *PostRepository) GetByID(id uint) (*entity.Post, error) {
	var post entity.Post
	err := r.db.Preload("User").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// GetByUserID 根据用户ID获取文章列表
func (r *PostRepository) GetByUserID(userID uint, offset, limit int) ([]entity.Post, int64, error) {
	var posts []entity.Post
	var total int64

	// 获取总数
	r.db.Model(&entity.Post{}).Where("user_id = ?", userID).Count(&total)

	// 获取文章列表
	err := r.db.Preload("User").Where("user_id = ?", userID).
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&posts).Error

	return posts, total, err
}

// GetAll 获取所有文章列表
func (r *PostRepository) GetAll(offset, limit int) ([]entity.Post, int64, error) {
	var posts []entity.Post
	var total int64

	// 获取总数
	r.db.Model(&entity.Post{}).Count(&total)

	// 获取文章列表
	err := r.db.Preload("User").
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&posts).Error

	return posts, total, err
}

// Update 更新文章
func (r *PostRepository) Update(post *entity.Post) error {
	return r.db.Save(post).Error
}

// Delete 删除文章
func (r *PostRepository) Delete(id uint, userID uint) (int64, error) {
	result := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&entity.Post{})
	return result.RowsAffected, result.Error
}

// UpdateByIdAndUserId 根据ID和用户ID更新文章
func (r *PostRepository) UpdateByIdAndUserId(id uint, userID uint, title, content string) (int64, error) {
	result := r.db.Model(&entity.Post{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(map[string]interface{}{
			"title":   title,
			"content": content,
		})
	return result.RowsAffected, result.Error
}
