package main

import (
	cfg "awesomeProject5/config"
	"awesomeProject5/internal/app"
	"awesomeProject5/internal/app/delivery"
	"awesomeProject5/internal/repository"
	"awesomeProject5/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg.InitConfig()

	engine := gin.Default()

	repo := repository.New()
	uc := usecase.New(repo)
	handlers := delivery.New(uc)

	app.InitEndPoints(engine, handlers)

	engine.Run(":8081") // http://localhost:8081
}
