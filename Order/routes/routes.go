package routes

import (
	"Order/controller"
	"Order/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 用户注册,登录
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// 商家注册,登录
	r.POST("/merchant/register", controller.MerchantRegister)
	r.POST("/merchant/login", controller.MerchantLogin)

	// 订单管理
	orderGroup := r.Group("/orders")
	orderGroup.Use(middleware.AuthMiddleware())
	orderGroup.POST("/Place", controller.PlaceOrder)
	orderGroup.POST("/Modify/:id", controller.ModifyOrder)
	//orderGroup.POST("/Cancel/:id", controller.CancelOrder)
	orderGroup.DELETE("/Cancel/:id", controller.CancelOrder)
	orderGroup.GET("/View", controller.ViewOrders)

	// 菜品管理
	dishGroup := r.Group("/dishes")
	dishGroup.POST("/Up", controller.UploadDish)
	dishGroup.GET("/Get", controller.GetDishes)

}
