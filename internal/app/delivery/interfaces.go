package delivery

import (
	"awesomeProject5/internal/entity"
	"context"
	"github.com/google/uuid"
)

type IUseCase interface {
	CreateTask(context.Context, entity.Task) (uuid.UUID, error)
}
