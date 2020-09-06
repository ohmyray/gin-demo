package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ohmyray/gin-demo01/common"
	"github.com/ohmyray/gin-demo01/model"
	"github.com/ohmyray/gin-demo01/uitl"
	"net/http"
)


func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	telephone := ctx.PostForm("telephone")

	// 判断参数是否符合需求
	if len(username) == 0 {
		username = uitl.RandomString(10)
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

	if isTelephoneExist(DB, telephone) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code":    20003,
				"message": "手机号已被注册，请勿重复注册！",
			})
		return
	}

	newUser := model.User{
		Username:  username,
		Password:  password,
		Telephone: telephone,
	}
	DB.Create(&newUser)

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":    1001,
			"message": "注册成功！",
		})
}


func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

