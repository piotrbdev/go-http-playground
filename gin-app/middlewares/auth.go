// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

//	func AuthMiddleware() gin.HandlerFunc {
//		return func(c *gin.Context) {
//			authHeader := c.GetHeader("Authorization")
//			if authHeader == "" {
//				c.JSON(http.StatusUnauthorized, models.Response{Message: "missing auth token"})
//				c.Abort()
//				return
//			}
//
//			if !strings.HasPrefix(authHeader, "Bearer ") {
//				c.JSON(http.StatusUnauthorized, models.Response{Message: "invalid auth format"})
//				c.Abort()
//				return
//			}
//
//			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
//
//			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
//				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//					return nil, fmt.Errorf("unexpected signing method")
//				}
//				return jwtSecret, nil
//			}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
//
//			if err != nil || !token.Valid {
//				c.JSON(http.StatusUnauthorized, models.Response{Message: "invalid token"})
//				c.Abort()
//				return
//			}
//			claims, ok := token.Claims.(jwt.MapClaims)
//			if !ok {
//				c.JSON(http.StatusUnauthorized, models.Response{Message: "invalid claims"})
//				c.Abort()
//				return
//			}
//
//			userIDFloat, ok := claims["user_id"].(float64)
//			if !ok {
//				c.JSON(http.StatusUnauthorized, models.Response{Message: "invalid user_id"})
//				c.Abort()
//				return
//			}
//
//			c.Set("userID", userIDFloat)
//
//			c.Next()
//		}
//	}
package middlewares

import (
	"gin-app/models"
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var keyFunc jwt.Keyfunc

func InitJWKS() error {
	k, err := keyfunc.NewDefault(
		[]string{
			"http://localhost:8081/realms/my-app/protocol/openid-connect/certs",
		},
	)
	if err != nil {
		return err
	}

	keyFunc = k.Keyfunc

	return nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.Response{
				Message: "missing token",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, models.Response{Message: "invalid auth format"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, keyFunc)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, models.Response{
				Message: "invalid token",
			})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID := claims["sub"].(string)

		c.Set("userID", userID)

		c.Next()
	}
}
