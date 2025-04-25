package controller

import (
	"Order/models"
	"Order/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 上传菜品
func UploadDish(c *gin.Context) {
	var dish models.Dish
	if err := c.ShouldBindJSON(&dish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	utils.DB.Create(&dish)
	c.JSON(http.StatusOK, gin.H{"message": "菜品上传成功"})
}

// 展示菜品
func GetDishes(c *gin.Context) {
	var dishes []models.Dish
	utils.DB.Find(&dishes)
	c.JSON(http.StatusOK, gin.H{"dishes": dishes})
}
