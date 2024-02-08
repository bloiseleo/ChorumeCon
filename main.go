package main

import (
	"github.com/bloiseleo/chorumecon/env"
	"github.com/bloiseleo/chorumecon/middlewares"
	v1 "github.com/bloiseleo/chorumecon/routes/v1"
	"github.com/gin-gonic/gin"
)

func setUpV1(e *gin.Engine) {
	v1Router := e.Group("v1")
	v1.Setup(v1Router)
}

func setupApplication() *gin.Engine {
	g := gin.New()
	g.Use(gin.LoggerWithFormatter(env.ChorumeconLogger))
	g.Use(middlewares.ValidateBodyForPOSTRequest())
	g.Use(gin.Recovery())
	setUpV1(g)
	return g
}

func main() {
	engine := setupApplication()
	engine.Run(":" + env.Env("port", "8080"))
}
