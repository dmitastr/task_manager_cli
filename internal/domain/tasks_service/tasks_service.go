package tasksservice

import (
	"slices"

	"github.com/dmastr/task-manager-cli/internal/common/datetimeformattter"
	"github.com/dmastr/task-manager-cli/internal/database/storage"
	"github.com/dmastr/task-manager-cli/internal/domain/entity"
)

type TasksService struct {
	db storage.Storage
}

func NewService(storage storage.Storage) *TasksService {
	return &TasksService{db: storage}
}

func (s *TasksService) Get(ID entity.TaskID) entity.Task {
	return s.db.Get(ID)
}

func (s *TasksService) GetAll() []entity.Task {
	tasks := s.db.GetAll()
	slices.SortFunc(tasks, func(a, b entity.Task) int {
		if a.IsCompleted == b.IsCompleted {
			if a.Text > b.Text {
				return 1
			}
			return -1
		} else if a.IsCompleted {
			return 1
		} 
		return -1
	})
	return tasks
}

func (s *TasksService) Put(title string) (entity.Task, error) {
	task := s.db.Put(title)
	return task, nil
}

func (s *TasksService) MarkTask(id string, isCompleted bool) (t entity.Task, err error) {
	idParsed, err := entity.IdFromString(id)
	if err != nil {
		return t, err
	}

	task := s.db.Get(idParsed)

	task.IsCompleted = isCompleted
	task.UpdatedAt = datetimeformattter.DateTimeNow()
	s.db.Update(task)
	return task, err
}
