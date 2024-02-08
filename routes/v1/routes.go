package v1

import (
	"github.com/bloiseleo/chorumecon/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup(v *gin.RouterGroup) {
	v.POST("/login", login)
	v.POST("/bonify", middlewares.JWTMiddleware(), bonify)
}
