package middleware

import (
	"Order/models"
	"Order/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginUser, err := c.Cookie("login_user")
		if err != nil || loginUser == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
			c.Abort()
			return
		}

		// 查用户ID
		var user models.User
		if err := utils.DB.Where("username = ?", loginUser).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
			c.Abort()
			return
		}

		//c.Set("user", loginUser)
		c.Set("user_id", strconv.Itoa(int(user.ID))) // 设置 user_id 给后面的handler用
		c.Next()
	}
}
