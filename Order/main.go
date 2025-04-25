package main

import (
	"Order/routes"
	"Order/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	//r := gin.Default()
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Hello, Gin!",
	//	})
	//})
	//
	//r.Run(":8080")

	//dsn := "root:050306@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	//_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal("Failed to connect to database:", err)
	//}
	//
	//log.Println("Database connected successfully!")

	utils.InitDB()
	//utils.InitRedis()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
