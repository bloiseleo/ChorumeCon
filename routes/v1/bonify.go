package v1

import (
	"errors"
	"net/http"
	"github.com/bloiseleo/chorumecon/database/repository"
	"github.com/bloiseleo/chorumecon/helpers"
	"github.com/bloiseleo/chorumecon/services/bonification"
	"github.com/gin-gonic/gin"
)

type bonifyDto struct {
	Coins     int    `json:"coins" binding:"required"`
	DiscordId string `json:"discordId" binding:"required"`
}

func extractUser(ctx *gin.Context) (string, error) {
	bonifyUserPossible, exists := ctx.Get("user")
	if !exists {
		return "", errors.New("not found")
	}
	u, ok := bonifyUserPossible.(string)
	if !ok {
		return "", errors.New("no compatible")
	}
	return u, nil
}

func extractExchange(ctx *gin.Context) (float32, error) {
	exchangePossible, exists := ctx.Get("exchange")
	if !exists {
		return 0.0, errors.New("exchange not found")
	}
	e, ok := exchangePossible.(float32)
	if !ok {
		return 0.0, errors.New("no compatible")
	}
	return e, nil
}

func bonify(ctx *gin.Context) {
	var dto bonifyDto
	bonifyUser, errEx := extractUser(ctx)
	if errEx != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}
	bonifyExchange, errEx := extractExchange(ctx)
	if errEx != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}
	if err := ctx.BindJSON(&dto); err != nil {
		error := helpers.ValidationErrorMessage(err, http.StatusUnprocessableEntity)
		ctx.JSON(error.Status, error)
		return
	}
	user := repository.FindUserByDiscordId(dto.DiscordId)
	if user == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "User of ID " + string(dto.DiscordId) + " was not found",
		})
		return
	}
	if dto.Coins < 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "You cannot remove coins from user",
		})
		return
	}
	coins := bonification.CalculateBonification(dto.Coins, bonifyExchange)
	repository.IncrementCoins(user, coins, bonifyUser)
	ctx.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"chorumecoins": coins,
	})
	return
}
