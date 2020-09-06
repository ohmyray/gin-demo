package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ohmyray/gin-demo01/model"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error)  {
	expirationTime := time.Now().Add(7*24*time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "ohmyray.top",
			Subject: "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
	//head 协议 SigningMethodHS256 eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
	// claims 信息 eyJVc2VySWQiOjUsImV4cCI6MTU5OTk3MDQ3NCwiaWF0IjoxNTk5MzY1Njc0LCJpc3MiOiJvaG15cmF5LnRvcCIsInN1YiI6InVzZXIgdG9rZW4ifQ
	// hash head + claims cqUUYg83mKfKbyEmiLFzFjrPNn_Mz6wz98btOZpM6bA
	//"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjUsImV4cCI6MTU5OTk3MDQ3NCwiaWF0IjoxNTk5MzY1Njc0LCJpc3MiOiJvaG15cmF5LnRvcCIsInN1YiI6InVzZXIgdG9rZW4ifQ.cqUUYg83mKfKbyEmiLFzFjrPNn_Mz6wz98btOZpM6bA"
}
