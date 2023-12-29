package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lyhgo/gin/tool/logging"
)

func JwtChecker() gin.HandlerFunc {
	return jwtCheckerHandler()
}

// jwtCheckerHandler returns a middleware to validate jwt token.
func jwtCheckerHandler() gin.HandlerFunc {
	logger := logging.GetLogger()
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		var accessTokenString string
		if len(authorization) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "empty token",
			})
			c.Abort()
			return
		}

		accessTokenString = strings.Split(authorization, " ")[1]
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "access token can not be validated",
					"err":     fmt.Sprintf("%v", err),
				})
				c.Abort()
			}
		}()
		ok, err := validateToken(accessTokenString, os.Getenv("SIGN_KEY_PATH"))
		switch {
		case err != nil:
			logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "access token can not be validated",
				"err":     err.Error(),
			})
			c.Abort()
			return
		case !ok:
			logger.Info("access token can not be validated")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "access token can not be validated",
				"err":     err.Error(),
			})
			c.Abort()
			return
		default:
			c.Next()
		}
	}
}

// validateToken validate token string
func validateToken(tokenString string, signKeyPath string) (bool, error) {
	signBytes, err := os.ReadFile(signKeyPath)
	if err != nil {
		return false, fmt.Errorf("load private rsa key failed")
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return false, fmt.Errorf("parse rsa key failed")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &signKey.PublicKey, nil
	})
	return token.Valid, err
}
