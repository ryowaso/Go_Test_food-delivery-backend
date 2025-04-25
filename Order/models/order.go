package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint    `gorm:"not null" json:"user_id"`
	DishID     uint    `gorm:"not null" json:"dish_id"`
	DishName   string  `gorm:"size:255;not null" json:"dish_name"`
	Quantity   uint    `gorm:"not null" json:"quantity"`
	TotalPrice float64 `gorm:"not null" json:"total_price"`
	Status     string  `json:"status"`
}
