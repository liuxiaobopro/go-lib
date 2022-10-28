package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/go-lib/ecode"
	"github.com/liuxiaobopro/go-lib/response"
)

var JwtKey = []byte("1dsasadsadasd")
var code int

type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// SetToken 生成token
func GetToken(username string) (string, ecode.BizErr) {
	expireTime := time.Now().Add(time.Hour * 10)
	SetClaims := JwtClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", ecode.ERROR
	}
	return token, ecode.SUCCSESS
}

// CheckToken 验证token
func CheckToken(token string) (*JwtClaims, ecode.BizErr) {
	setToken, _ := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*JwtClaims); setToken.Valid {
		return key, ecode.SUCCSESS
	} else {
		return nil, ecode.ERROR
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		if tokenHerder == "" {
			c.JSON(http.StatusOK, response.GetErrRes(ecode.ERROR_TOKEN_INVALID))
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, response.GetErrRes(ecode.ERROR_TOKEN_INVALID))
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode.Code != 0 {
			c.JSON(http.StatusOK, response.GetErrRes(ecode.ERROR_TOKEN_INVALID))
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			c.JSON(http.StatusOK, response.GetErrRes(ecode.ERROR_TOKEN_INVALID))
			c.Abort()
			return
		}
		// c.Set("username", key.Username)
		c.Next()

	}
}
