package storage

import (

	"github.com/dmastr/task-manager-cli/internal/domain/entity"
)


type Storage interface {
	Get(ID entity.TaskID) entity.Task
	GetAll() []entity.Task
	Put(string) entity.Task
	Update(entity.Task)
	// Delete(ID string)
}