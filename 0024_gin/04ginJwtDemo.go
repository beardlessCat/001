package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

const TokenExpireDuration = time.Hour * 24

var CustomSecret = []byte("夏天夏天悄悄过去")

type CustomClaims struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

func generalToken() (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		"lyj", // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "my-project", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

/*
*认证获取jwt
 */
func authHandler(c *gin.Context) {

}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    9999,
				"message": "token must not be null!",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 9999,
				"msg":  "token format err!",
			})
			c.Abort()
			return
		}
		customClaims, err := ParseToken(parts[1])
		//token解析失败
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    9999,
				"message": "invalidate token!",
			})
			c.Abort()
			return
		}
		username := customClaims.Username
		c.Set("username", username)
		c.Next()
	}
}
func main() {
	engine := gin.Default()
	engine.POST("/auth", func(context *gin.Context) {
		token, _ := generalToken()
		context.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	engine.GET("/getUser", JWTAuthMiddleware(), func(context *gin.Context) {
		username, _ := context.Get("username")
		context.JSON(http.StatusOK, gin.H{
			"username": username,
		})
	})
	engine.Run()
}
