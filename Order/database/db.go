package database

//import (
//	"Order/config"
//	"Order/models"
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//)
//
//var DB *gorm.DB
//
//func InitDB() {
//
//	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//	//	config.DBUser, config.DBPwd, config.DBHost, config.DBPort, config.DBName)
//	//
//	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	//if err != nil {
//	//	panic("数据库连接失败")
//	//}
//
//	var db, err = gorm.Open(mysql.Open(config.MySOLDSN), &gorm.Config{})
//	if err != nil {
//		log.Fatal("数据库连接失败:", err)
//	}
//
//	db.AutoMigrate(&models.User{}, &models.Order{}, &models.Dish{}, &models.Merchant{})
//	DB = db
//	fmt.Println("数据库连接成功")
//}
