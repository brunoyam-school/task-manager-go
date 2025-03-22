package app

import (
	"awesomeProject5/internal/app/delivery"
	"awesomeProject5/internal/repository"
	"awesomeProject5/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
}

type App struct {
	engine *gin.Engine
	logger ILogger
}

func New(logger ILogger) *App {
	engine := gin.Default()
	handlers := delivery.New(usecase.New(repository.New()))

	InitEndPoints(engine, handlers)

	return &App{
		engine: engine,
		logger: logger}
}

func (a *App) Run() error {
	return a.engine.Run(viper.GetString("port")) // http://localhost:8081
}
