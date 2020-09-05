package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 全局前缀
const apiPreFix = "/api/v1"

// 模型
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null"`
	Password  string `gorm:"size:255;not null"`
	Telephone string `gorm:"type:varchar(11);not null"`
}

func main() {
	// 获取数据库连接
	db := InitDB()
	defer db.Close()

	r := gin.Default()

	r.POST(apiPreFix+"/register", func(ctx *gin.Context) {

		// 获取参数
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		telephone := ctx.PostForm("telephone")

		// 判断参数是否符合需求
		if len(username) == 0 {
			username = RandomString(10)
		}

		if len(password) < 6 {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"code":    20001,
					"message": "密码长度不够！",
				})
			return
		}

		if len(telephone) != 11 {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"code":    20002,
					"message": "手机号长度不正确！",
				})
			return
		}

		if isTelephoneExist(db, telephone) {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"code":    20003,
					"message": "手机号已被注册，请勿重复注册！",
				})
			return
		}

		newUser := User{
			Username:  username,
			Password:  password,
			Telephone: telephone,
		}
		db.Create(&newUser)

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code":    1001,
				"message": "注册成功！",
			})
	})

	r.Run()
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(n int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

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

	db.AutoMigrate(&User{})

	return db
}
