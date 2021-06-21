package handler

import (
	"server-api/auth"
	"server-api/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
		c.Next()
	}
}

func Middleware(authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || len(authHeader) == 0 {
			errorResponse := helper.FailResponse(400, "unauthorized user", gin.H{"error": "unauthorized user"})
			c.AbortWithStatusJSON(400, errorResponse)
		}
		token, err := authService.ValidateToken(authHeader)
		if err != nil {
			errorResponse := helper.FailResponse(400, "unauthorized user", gin.H{"error": "unauthorized user"})
			c.AbortWithStatusJSON(400, errorResponse)
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		print(claim)
		if !ok {
			errorResponse := helper.FailResponse(400, "unauthorized user", gin.H{"error": "unauthorized user"})
			c.AbortWithStatusJSON(400, errorResponse)
		}
		ID := int(claim["user_id"].(float64))
		c.Set("currentUser", ID)
	}
}
