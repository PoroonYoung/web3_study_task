package dto

// CreatePostRequest 创建文章请求
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
}

// PostResponse 文章响应
type PostResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

// PostListResponse 文章列表响应
type PostListResponse struct {
	Posts []PostResponse `json:"posts"`
	Total int64          `json:"total"`
}

type UpdatePostRequest struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
	UserID  uint   `json:"userId"`
}
