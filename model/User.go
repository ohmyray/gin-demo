package model

import "github.com/jinzhu/gorm"

// 模型
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(18);not null"`
	Password  string `gorm:"size:255;not null"`
	Telephone string `gorm:"type:varchar(11);not null"`
}
