package usecase

import (
	"awesomeProject5/internal/entity"
	"github.com/google/uuid"
)

type UseCase struct {
	repository IDatabase
}

func New(repo IDatabase) *UseCase {
	return &UseCase{repository: repo}
}

func (uc *UseCase) CreateTask(task entity.Task) (uuid.UUID, error) {
	task.ID = uuid.New()

	return task.ID, uc.repository.Create(task)
}
