package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"web3_study_task/task4/pkg/response"
	"web3_study_task/task4/pkg/utils"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware(jwtUtil *utils.Util) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			response.Unauthorized(c, "Authorization header 缺失")
			c.Abort()
			return
		}

		token := strings.TrimPrefix(bearerToken, "Bearer ")
		if token == bearerToken {
			response.Unauthorized(c, "Authorization header 格式错误，应为 'Bearer <token>'")
			c.Abort()
			return
		}

		// 验证token
		claims, err := jwtUtil.Verify(token)
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("token无效: %v", err))
			c.Abort()
			return
		}

		// 将用户ID存储到上下文中
		if claims.Subject != "" {
			var userID uint
			fmt.Sscanf(claims.Subject, "%d", &userID)
			c.Set("userID", userID)
			c.Set("username", claims.Data["username"])
		}

		c.Next()
	}
}

// GetCurrentUserID 从上下文中获取当前用户ID
func GetCurrentUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, false
	}

	id, ok := userID.(uint)
	return id, ok
}
