package dto

type CommentPublishRequest struct {
	PostID  uint `json:"postId" binding:"required"`
	UserID  uint
	Content string `json:"content" binding:"required,min=1"`
}
