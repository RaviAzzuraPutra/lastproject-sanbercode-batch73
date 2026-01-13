package middleware

import (
	"last-project/app/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddlewar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			err := helper.NewUnauthorizedError("Unauthorized")
			if appErr, ok := err.(*helper.AppError); ok {
				ctx.JSON(appErr.Code, gin.H{
					"Message": appErr.Message,
				})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Internal Server Error",
			})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			err := helper.NewUnauthorizedError("Invalid Format")
			if appErr, ok := err.(*helper.AppError); ok {
				ctx.JSON(appErr.Code, gin.H{
					"Message": appErr.Message,
				})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Internal Server Error",
			})
			ctx.Abort()
			return
		}

		tokenString := parts[1]

		claims, errClaims := helper.ParseJWT(tokenString)

		if errClaims != nil {
			if appErr, ok := errClaims.(*helper.AppError); ok {
				ctx.JSON(appErr.Code, gin.H{
					"Message": appErr.Message,
				})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token is invalid or has expired.",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)

		ctx.Next()
	}
}
