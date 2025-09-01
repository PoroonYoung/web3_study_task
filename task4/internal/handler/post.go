package handler

import (
	"strconv"
	"web3_study_task/task4/internal/middleware"
	"web3_study_task/task4/internal/service"
	"web3_study_task/task4/pkg/dto"
	"web3_study_task/task4/pkg/response"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

// CreatePost 创建文章
func (h *PostHandler) CreatePost(c *gin.Context) {
	// 获取当前用户ID
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.Unauthorized(c, "无法获取用户信息")
		return
	}

	// 绑定请求参数
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数格式错误: "+err.Error())
		return
	}

	// 调用服务层创建文章
	postResp, err := h.postService.CreatePost(userID, &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, postResp)
}

// GetPost 获取文章详情
func (h *PostHandler) GetPost(c *gin.Context) {
	// 获取文章ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "文章ID格式错误")
		return
	}

	// 调用服务层获取文章
	postResp, err := h.postService.GetPost(uint(id))
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, postResp)
}

// GetUserPosts 获取用户的文章列表
func (h *PostHandler) GetUserPosts(c *gin.Context) {
	// 获取当前用户ID
	userID, ok := middleware.GetCurrentUserID(c)
	if !ok {
		response.Unauthorized(c, "无法获取用户信息")
		return
	}

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 调用服务层获取文章列表
	listResp, err := h.postService.GetUserPosts(userID, page, pageSize)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, listResp)
}

// GetAllPosts 获取所有文章列表
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 调用服务层获取文章列表
	listResp, err := h.postService.GetAllPosts(page, pageSize)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, listResp)
}

func (h *PostHandler) UpdatePost(context *gin.Context) {
	var req dto.UpdatePostRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		response.BadRequest(context, "参数格式错误: "+err.Error())
		return
	}
	userID, _ := middleware.GetCurrentUserID(context)
	req.UserID = userID
	_, err := h.postService.UpdatePost(req)
	if err != nil {
		response.InternalServerError(context, err.Error())
		return
	}
	response.Success(context, nil)
}

func (h *PostHandler) DeletePost(context *gin.Context) {
	var idStr string
	idStr = context.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(context, "ID格式不合法")
	}
	//获取当前登录用户id
	userID, _ := middleware.GetCurrentUserID(context)
	e := h.postService.DeletePost(uint(id), userID)
	if e != nil {
		response.InternalServerError(context, e.Error())
		return
	}
	response.Success(context, nil)
}
