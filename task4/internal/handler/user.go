package handler

import (
	"web3_study_task/task4/internal/middleware"
	"web3_study_task/task4/internal/service"
	"web3_study_task/task4/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if username == "" || password == "" {
		response.BadRequest(c, "用户名和密码不能为空")
		return
	}

	loginResp, err := h.userService.Login(username, password)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, loginResp)
}

// GetUserInfo 获取用户信息
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.Unauthorized(c, "无法获取用户信息")
		return
	}

	userInfo, err := h.userService.GetUserInfo(userID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, userInfo)
}

// Hello 测试接口
func (h *UserHandler) Hello(c *gin.Context) {
	panic("异常处理测试")
	response.Success(c, "hello world!")
}
