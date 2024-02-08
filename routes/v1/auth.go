package v1

import (
	"net/http"

	"github.com/bloiseleo/chorumecon/database/repository"
	"github.com/bloiseleo/chorumecon/helpers"
	"github.com/bloiseleo/chorumecon/services/auth"
	"github.com/gin-gonic/gin"
)

type loginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func login(ctx *gin.Context) {
	var dto loginDTO
	if err := ctx.BindJSON(&dto); err != nil {
		error := helpers.ValidationErrorMessage(err, http.StatusUnprocessableEntity)
		ctx.JSON(error.Status, error)
		return
	}
	user := repository.FindApiUserByName(dto.Username)
	if user == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Wrong Credentials",
		})
		return
	}
	if !user.ComparePassword(dto.Password) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": "Forbbiden",
		})
		return
	}
	token := auth.CreateToken(user.Id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"token":   token,
		"message": "Authenticated",
	})
	return
}
