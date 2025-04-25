package controller

import (
	"Order/models"
	"Order/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 下单
func PlaceOrder(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("user_id"))

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	order.UserID = uint(userID)
	order.Status = "待解决"

	var dish models.Dish
	//if err := utils.DB.First(&dish, order.DishID).Error; err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "菜品不存在"})
	//	return
	//}
	if err := utils.DB.Where("name = ?", order.DishName).First(&dish).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "菜品不存在"})
		return
	}

	order.TotalPrice = float64(order.Quantity) * dish.Price

	if err := utils.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "订单创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "下单成功", "order": order})
}

// 修改订单
//func ModifyOrder(c *gin.Context) {
//	orderID := c.Param("id")
//
//	var order models.Order
//	if err := utils.DB.First(&order, orderID).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "订单不存在"})
//		return
//	}
//
//	// 只能修改未完成的订单
//	if order.Status != "待解决" {
//		c.JSON(http.StatusForbidden, gin.H{"error": "订单无法修改"})
//		return
//	}
//
//	if err := c.ShouldBindJSON(&order); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
//		return
//	}
//
//	utils.DB.Save(&order)
//	c.JSON(http.StatusOK, gin.H{"message": "订单修改成功", "order": order})
//}

func ModifyOrder(c *gin.Context) {
	orderID := c.Param("id")

	var existingOrder models.Order
	if err := utils.DB.First(&existingOrder, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "订单不存在"})
		return
	}

	if existingOrder.Status != "待解决" {
		c.JSON(http.StatusForbidden, gin.H{"error": "订单无法修改"})
		return
	}

	//更新字段
	var updateData struct {
		DishID   uint   `json:"dish_id"`
		Quantity uint   `json:"quantity"`
		Status   string `json:"status"`

		DishName string `json:"dish_name"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 更新字段
	existingOrder.DishID = updateData.DishID
	existingOrder.Quantity = updateData.Quantity
	existingOrder.Status = updateData.Status

	existingOrder.DishName = updateData.DishName

	//查找价格
	var dish models.Dish
	//if err := utils.DB.First(&dish, existingOrder.DishID).Error; err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "菜品不存在"})
	//	return
	//}
	if err := utils.DB.Where("name = ?", updateData.DishName).First(&dish).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "订单保存失败"})
		return
	}
	//重新计算
	existingOrder.TotalPrice = float64(existingOrder.Quantity) * dish.Price

	utils.DB.Save(&existingOrder)

	c.JSON(http.StatusOK, gin.H{
		"message": "订单修改成功",
		"order":   existingOrder,
	})
}

// 查看订单
func ViewOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10
	offset := (page - 1) * limit

	var orders []models.Order
	utils.DB.Offset(offset).Limit(limit).Find(&orders)

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// 取消订单
func CancelOrder(c *gin.Context) {
	orderID := c.Param("id")

	var order models.Order
	if err := utils.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "订单不存在"})
		return
	}

	if order.Status != "待解决" {
		c.JSON(http.StatusForbidden, gin.H{"error": "订单无法取消"})
		return
	}

	//order.Status = "已取消"
	//utils.DB.Save(&order)

	if err := utils.DB.Unscoped().Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "订单删除失败"})
		return
	} //hard

	c.JSON(http.StatusOK, gin.H{"message": "订单取消成功"})
}
