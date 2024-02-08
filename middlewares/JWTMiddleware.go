package middlewares

import (
	"net/http"
	"strings"
	"github.com/bloiseleo/chorumecon/database/entity"
	"github.com/bloiseleo/chorumecon/services/auth"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(header, "Bearer ")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}
		token, err := auth.ValidateJWT(tokenString)
		if err != nil {
			if auth.TokenExpired(err) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"message": "Unauthorized",
				})
				return
			}
			panic(err)
		}
		e := entity.AdaptFromToken(token)
		ctx.Set("user", e.Username)
		ctx.Set("exchange", e.Exchange)
		ctx.Next()
		return
	}
}
