package usecase

import "awesomeProject5/internal/entity"

type IDatabase interface {
	Create(task entity.Task) error
}
