package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ohmyray/gin-demo01/common"
	"github.com/ohmyray/gin-demo01/model"
	"github.com/ohmyray/gin-demo01/uitl"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(ctx *gin.Context) {
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2004,
			"message": "用户未注册！",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2005,
			"message": "密码错误!",
		})
		return
	}

	token := "111"

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1001,
		"message": "登录成功！",
		"data": gin.H{
			"token": token,
		},
	})

}

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	telephone := ctx.PostForm("telephone")

	// 判断参数是否符合需求
	log.Println("username", username)
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

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(
		http.StatusInternalServerError,
			gin.H{
				"code":    40001,
				"message": "系统错误！" + err.Error(),
			})
		return
	}

	newUser := model.User{
		Username:  username,
		Password:  string(hasedPassword),
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
