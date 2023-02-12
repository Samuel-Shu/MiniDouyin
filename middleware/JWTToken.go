package middleware

import (
	"MiniDouyin/utils"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var(
	TokenISNotCorrect error = errors.New("Token不正确")
	TokenExpired error = errors.New("Token超时，已过期")
)

var MySigningKey = []byte("fjonhgelnvlsalpq")

func BuildToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"name":"shuxin",
		"exp":time.Now().Unix()+60*10,
	})
	tokenString ,err := token.SignedString(MySigningKey)
	utils.ResolveError(err)
	fmt.Println(tokenString)
	return tokenString
}

func ParseToken(tokenString string)  {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return MySigningKey, nil
	})
	utils.ResolveError(err)
	if err == TokenISNotCorrect{
		fmt.Println(TokenISNotCorrect)
	}
	fmt.Println("token:", token)
	fmt.Println("token:", token.Valid)
}