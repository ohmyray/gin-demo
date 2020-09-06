package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ohmyray/gin-demo01/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "gindemo01"
	charset := "utf8"
	username := "root"
	password := "root"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)

	// 开启数据库连接
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}