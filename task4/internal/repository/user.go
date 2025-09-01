package repository

import (
	"gorm.io/gorm"
	"web3_study_task/task4/pkg/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetByCredentials 根据用户名和密码获取用户
func (r *UserRepository) GetByCredentials(username, password string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID 根据ID获取用户
func (r *UserRepository) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create 创建用户
func (r *UserRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

// Update 更新用户
func (r *UserRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
