package service

import (
	"github.com/dmastr/task-manager-cli/internal/domain/entity"
)


type Service interface {
	GetAll() []entity.Task
	Get(id entity.TaskID) entity.Task
	Put(title string) (entity.Task, error)
	MarkTask(id string, isCompleted bool) (entity.Task, error)
} 