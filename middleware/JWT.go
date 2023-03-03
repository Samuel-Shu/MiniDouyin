package middleware

import (
	"MiniDouyin/config"
	"MiniDouyin/model"
	"MiniDouyin/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type MyClaim struct {
	UserId   int32
	UserName string
	jwt.RegisteredClaims
}

// GenerateToken Generate generate jwtToken
func GenerateToken(uerName string, userId int32) string {
	claim := MyClaim{
		UserId:   userId,
		UserName: uerName,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                    //effective time
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    //sign time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)), //expire time
		},
	}
	//use HAS256 to sign a jwtToken with claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, _ := token.SignedString(config.JwtKey)
	return jwtToken
}

// ParseToken  can parse jwtToken to get Claim's information
func ParseToken(jwtToken string) (interface{}, error) {
	claims := MyClaim{}
	if jwtToken == "" {
		return nil, errors.New("you don't have valid token")
	}
	_, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("your alg is not correct,your alg is %v", token.Header["alg"])
		}
		return config.JwtKey, nil
	})
	if err != nil {
		return claims, err
	}
	return claims, nil
}

//JWT gin中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		type TokenData struct {
			Token string `json:"token" form:"token" binding:"required"`
		}
		var tokenData TokenData
		err := c.ShouldBind(&tokenData)
		if err != nil {
			fmt.Println(err)
		}
		if tokenData.Token == "" {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: utils.FAIL,
				StatusMsg:  utils.GetStatusMsg(utils.ERROR_TOKEN_EXIST),
			},
			)
			c.Abort()
			return
		}

		parseToken, err1 := ParseToken(tokenData.Token)

		if err1 != nil {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: utils.FAIL,
				StatusMsg:  utils.GetStatusMsg(utils.ERROR_TOKEN_WRONG),
			})
			c.Abort()
			return
		}

		if !model.FindUserWithId(parseToken.(MyClaim).UserId) {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: utils.FAIL,
				StatusMsg:  utils.GetStatusMsg(utils.ERROR_TOKEN_WRONG),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
