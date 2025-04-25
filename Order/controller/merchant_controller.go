package controller

import (
	"Order/models"
	"Order/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 商家注册
func MerchantRegister(c *gin.Context) {
	var merchant models.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	merchant.Password = utils.HashPassword(merchant.Password)

	if err := utils.DB.Create(&merchant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "商家注册成功"})
}

// 商家登录
func MerchantLogin(c *gin.Context) {
	var merchant models.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	var dbMerchant models.Merchant
	if err := utils.DB.Where("name = ?", merchant.Name).First(&dbMerchant).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "商家名或密码错误"})
		return
	}

	if !utils.CheckPassword(merchant.Password, dbMerchant.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "商家名或密码错误"})
		return
	}

	c.SetCookie("merchant_user", merchant.Name, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "商家登录成功"})
}
