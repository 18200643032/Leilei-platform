package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"context"
	e2 "Leilei-platform/public/e"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			e2.ErrorMsg(c, e2.StatusUnauthorized, e2.MsgFlags(e2.ERROR_AUTH))
			return
		}
		// 在这里执行验证token的逻辑
		valid, err := ValidateToken(token)
		if err != nil || !valid {
			e2.ErrorMsg(c, e2.StatusUnauthorized, e2.MsgFlags(e2.ERROR_AUTH))
			return
		}
		rdb := NewRedisClient()
		ctx := context.Background()
		err = rdb.Set(ctx, "token", token, 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器错误"})
			return
		}
		c.Next()
	}
}

func ValidateToken(tokenString string) (bool, error) {
	key := []byte("zhengzhong") // 密钥要与生成token时使用的一致

	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	return true, nil
}
