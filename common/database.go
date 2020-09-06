package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ohmyray/gin-demo01/model"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("dataSource.driverName")
	host := viper.GetString("dataSource.host")
	port := viper.GetString("dataSource.port")
	database := viper.GetString("dataSource.database")
	charset := viper.GetString("dataSource.charset")
	username := viper.GetString("dataSource.username")
	password := viper.GetString("dataSource.password")
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