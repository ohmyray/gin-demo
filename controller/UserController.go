package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ohmyray/gin-demo01/common"
	"github.com/ohmyray/gin-demo01/dto"
	"github.com/ohmyray/gin-demo01/model"
	"github.com/ohmyray/gin-demo01/response"
	"github.com/ohmyray/gin-demo01/uitl"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	if user == nil {
		response.Fail(ctx, nil, "系统错误!")
		return
	}
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "获取用户信息成功！")
}

func Login(ctx *gin.Context) {
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Success(ctx, nil, "用户未注册！")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Success(ctx, nil, "密码错误!")
		return
	}

	token, err := common.ReleaseToken(user)

	if err != nil {
		response.Fail(ctx, nil, "系统异常！")
		log.Printf("token generate error: %v", err)
		return
	}

	response.Success(ctx, gin.H{"token": token}, "登录成功！")

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
		response.Success(ctx, nil, "密码长度不够！")
		return
	}

	if len(telephone) != 11 {
		response.Success(ctx, nil, "手机号长度不正确！")
		return
	}

	if isTelephoneExist(DB, telephone) {
		response.Success(ctx, nil, "手机号已被注册，请勿重复注册！")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(ctx, nil, "系统错误,err: "+err.Error())
		return
	}

	newUser := model.User{
		Username:  username,
		Password:  string(hasedPassword),
		Telephone: telephone,
	}

	count := DB.Create(&newUser)
	if count.Error != nil {
		response.Fail(ctx, nil, "内部错误!")
		return
	}

	response.Success(ctx, nil, "注册成功！")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
