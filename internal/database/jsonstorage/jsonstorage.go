package jsonstorage

import (
	"encoding/json"
	"os"

	dttm "github.com/dmastr/task-manager-cli/internal/common/datetimeformattter"
	"github.com/dmastr/task-manager-cli/internal/domain/entity"
)

type DB struct {
	MetaInfo MetaInfo      `json:"meta_info"`
	Tasks    []entity.Task `json:"tasks"`
}

type MetaInfo struct {
	ItemsCount  int           `json:"items_count"`
	LastUpdated dttm.DateTime `json:"last_updated"`
	MaxId       int           `json:"max_id"`
}

type JsonStorage struct {
	dbPath string
	data   DB
}

func NewJsonStorage(dbPath string) *JsonStorage {
	js := JsonStorage{dbPath: dbPath}
	data, err := js.ReadData()
	if err != nil {
		panic(err)
	}
	js.data = data
	return &js
}

func (js *JsonStorage) GetAll() []entity.Task {
	data, err := js.ReadData()
	if err != nil {
		panic(err)
	}
	return data.Tasks
}

func (js *JsonStorage) Get(ID entity.TaskID) (t entity.Task) {
	data, err := js.ReadData()
	if err != nil {
		panic(err)
	}

	for _, task := range data.Tasks {
		if task.Id == ID {
			return task
		}
	}

	return t
}

func (js *JsonStorage) Put(title string) (t entity.Task) {
	task := entity.NewTask(title)
	task.Id = js.GetNextId()

	js.data.Tasks = append(js.data.Tasks, task)
	js.SaveData()
	return task
}


func (js *JsonStorage) Update(updatedTask entity.Task) {
	for i, t := range js.data.Tasks {
		if t.Id == updatedTask.Id {
			js.data.Tasks[i] = updatedTask
		}
	}
	js.SaveData()
}

func (js *JsonStorage) ReadData() (DB, error) {
	data := DB{}
	file, err := os.ReadFile(js.dbPath)
	if err != nil {
		return data, err
	}

	return data, json.Unmarshal(file, &data)
}

func (js *JsonStorage) GetNextId() entity.TaskID {
	id := js.GetMaxId() + 1
	return id
}

func (js *JsonStorage) GetMaxId() (maxId entity.TaskID) {
	for _, t := range js.data.Tasks {
		maxId = max(t.Id, maxId)
	}
	return
}

func (js *JsonStorage) SaveData() error {
	js.data.MetaInfo.LastUpdated = dttm.DateTimeNow()
	jsonData, err := json.MarshalIndent(js.data, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(js.dbPath, jsonData, 0644); err != nil {
		return err
	}
	return nil
}
