package service

import (
	"fmt"
	"time"
	"web3_study_task/task4/config"
	"web3_study_task/task4/internal/repository"
	"web3_study_task/task4/pkg/utils"
)

type UserService struct {
	userRepo *repository.UserRepository
	jwtUtil  *utils.Util
}

func NewUserService(userRepo *repository.UserRepository, jwtUtil *utils.Util) *UserService {
	return &UserService{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Token     string      `json:"token"`
	ExpiresIn int         `json:"expires_in"`
	User      interface{} `json:"user"`
}

// UserInfo 用户信息结构
type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Login 用户登录
func (s *UserService) Login(username, password string) (*LoginResponse, error) {
	// 验证用户凭据
	user, err := s.userRepo.GetByCredentials(username, password)
	if err != nil {
		return nil, fmt.Errorf("用户名或密码错误")
	}

	// 生成JWT token
	tokenString, err := s.jwtUtil.Sign(
		fmt.Sprintf("%d", user.ID),
		time.Hour*time.Duration(config.AppConfig.JWT.ExpireHours),
		map[string]any{
			"username": user.Username,
			"email":    user.Email,
			"role":     "user",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("生成token失败: %v", err)
	}

	return &LoginResponse{
		Token:     tokenString,
		ExpiresIn: config.AppConfig.JWT.ExpireHours * 3600,
		User: UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(userID uint) (*UserInfo, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}

	return &UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
