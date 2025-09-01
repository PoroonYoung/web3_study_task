package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web3_study_task/task4/internal/middleware"
	"web3_study_task/task4/internal/service"
	"web3_study_task/task4/pkg/dto"
	"web3_study_task/task4/pkg/entity"
	"web3_study_task/task4/pkg/response"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (h *CommentHandler) GetAllByPostId(context *gin.Context) {
	param := context.Param("postId")
	postId, err := strconv.Atoi(param)
	if err != nil {
		response.BadRequest(context, "文章id格式错误")
		return
	}
	//调用service处理
	var commentList *[]entity.Comment
	commentList, err = h.commentService.GetAllByPostId(uint(postId))
	if err != nil {
		response.InternalServerError(context, err.Error())
		return
	}
	response.Success(context, commentList)
}

func (h *CommentHandler) PublishComment(context *gin.Context) {
	var req dto.CommentPublishRequest
	err := context.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(context, "参数格式异常")
	}
	userId, _ := middleware.GetCurrentUserID(context)
	req.UserID = userId
	//调用service
	err = h.commentService.PublishComment(&req)
	if err != nil {
		response.InternalServerError(context, err.Error())
		return
	}
	response.Success(context, nil)
}
