package usecase

import (
	"awesomeProject5/internal/entity"
	"context"
	"github.com/google/uuid"
)

type UseCase struct {
	repository IDatabase
}

func New(repo IDatabase) *UseCase {
	return &UseCase{repository: repo}
}

func (uc *UseCase) CreateTask(ctx context.Context, task entity.Task) (uuid.UUID, error) {
	task.ID = uuid.New()

	return task.ID, uc.repository.Create(ctx, task)
}
