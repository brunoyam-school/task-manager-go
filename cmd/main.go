package main

import (
	cfg "awesomeProject5/config"
	"awesomeProject5/internal/app"
	"awesomeProject5/pkg/logger"
)

func main() {
	cfg.Initialize()
	appLogger := logger.Initialize()

	if err := app.New(appLogger).Run(); err != nil {
		return
	}
}
