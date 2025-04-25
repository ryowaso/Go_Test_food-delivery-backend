package models

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	ImageURL   string  `json:"image_url"`
	MerchantID string  `json:"merchant_id"`
}
