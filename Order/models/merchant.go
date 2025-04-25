package models

import (
	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement:false"`
	Name        string `gorm:"type:varchar(100);unique;not null"`
	ContactInfo string `gorm:"type:text"`
	Password    string `gorm:"type:varchar(100);not null"`
}
