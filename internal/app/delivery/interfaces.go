package delivery

import (
	"awesomeProject5/internal/entity"
	"github.com/google/uuid"
)

type IUseCase interface {
	CreateTask(task entity.Task) (uuid.UUID, error)
}
