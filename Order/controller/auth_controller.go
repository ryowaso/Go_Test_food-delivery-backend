package controller

import (
	"Order/models"
	"Order/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	user.Password = utils.HashPassword(user.Password)

	if err := utils.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// 用户登录
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	fmt.Println("前端传的用户名:", user.Username)
	fmt.Println("前端传的密码:", user.Password)

	var dbUser models.User
	//查数据库
	if err := utils.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	fmt.Println("数据库查到的用户名:", dbUser.Username)
	fmt.Println("数据库查到的密码（hash）:", dbUser.Password)
	//检验密码
	if !utils.CheckPassword(user.Password, dbUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}
	//设Cookie
	c.SetCookie("login_user", dbUser.Username, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "登陆成功"})
}
