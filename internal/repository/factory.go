package repository

import (
	"awesomeProject5/internal/entity"
	"awesomeProject5/internal/repository/sqlite"
	"context"
	"github.com/spf13/viper"
)

type IDatabase interface {
	Create(context.Context, entity.Task) error
}

func New() IDatabase {
	switch viper.GetString("db_type") {
	case "sqlite":
		return sqlite.New()
	case "postgres":

	}

	return nil
}
