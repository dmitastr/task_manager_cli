package memstorage

import (
	"sync"

	"github.com/dmastr/task-manager-cli/internal/domain/entity"
	"github.com/dmastr/task-manager-cli/internal/common/idgenerate"
)

type MemStorage struct {
	sync.Mutex
	idgenerate.IdGenerator
	tasks map[entity.TaskID]entity.Task
}

func NewMemStorage() *MemStorage {
	tasks := make(map[entity.TaskID]entity.Task)
	m := MemStorage{tasks: tasks}
	return &m
}

func (m *MemStorage) Get(ID entity.TaskID) (t entity.Task) {
	m.Lock()
	defer m.Unlock()
	t = m.tasks[ID]
	return t
}

func (m *MemStorage) GetAll() (tasks []entity.Task) {
	m.Lock()
	defer m.Unlock()
	for _, t := range m.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

func (m *MemStorage) Put(title string) entity.Task {
	task := entity.NewTask(title)
	task.Id = m.ID()
	m.tasks[task.Id] = task
	return task
}

func (m *MemStorage) Update(title string) {
	m.Put(title)
}

func (m *MemStorage) Delete(ID entity.TaskID) {
	delete(m.tasks, ID)
}
