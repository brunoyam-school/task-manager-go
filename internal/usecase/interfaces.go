package usecase

import (
	"awesomeProject5/internal/entity"
	"context"
)

type IDatabase interface {
	Create(context.Context, entity.Task) error
}
